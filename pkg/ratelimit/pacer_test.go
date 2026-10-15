package ratelimit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// A burst of 10 refilling at one request a second. The pacer holds one request back, so it will let 9 of those 10 go
// at once and then settle to a request a second.
const testPolicy = "l=10;rpm=60"

const (
	emissionInterval     = time.Second
	twoEmissionIntervals = 2 * time.Second
)

// newTestPacer returns a pacer on a clock that only moves when the test moves it, so the delays are exact.
func newTestPacer() (*pacer, *time.Duration) {
	now := new(time.Duration)
	*now = time.Hour // well clear of zero, so nothing is accidentally relying on the clock starting there
	return &pacer{now: func() time.Duration { return *now }}, now
}

func TestPacer_BeforeTheServerHasToldUsAnything_DoesNotWait(t *testing.T) {
	p, _ := newTestPacer()

	require.Zero(t, p.reserveSlot())
}

func TestPacer_WhenTheServerSendsNoRateLimitingHeaders_DoesNotWait(t *testing.T) {
	p, _ := newTestPacer()

	// A server with no policy for this caller says nothing at all, and we shouldn't invent a limit for it.
	p.observeResponse("", "")

	require.Zero(t, p.reserveSlot())
}

func TestPacer_WithQuotaToSpare_DoesNotWait(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=5;t=5")

	require.Zero(t, p.reserveSlot())
}

func TestPacer_WhenOnlyTheReservedRequestIsLeft_WaitsForOneRequestToComeBack(t *testing.T) {
	p, _ := newTestPacer()

	// One request left means the server would let us through, but the request after that would be rejected, and the
	// two sides' arithmetic doesn't agree to the last fraction of a request. So we stop here.
	p.observeResponse(testPolicy, "r=1;t=9")

	require.Equal(t, emissionInterval, p.reserveSlot())
}

func TestPacer_WhenTheQuotaIsSpent_WaitsForTheReservedRequestToComeBackToo(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=0;t=10")

	require.Equal(t, twoEmissionIntervals, p.reserveSlot())
}

func TestPacer_WithNoFurtherResponses_SpendsTheBurstAndThenPacesItself(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=10;t=0")

	// A burst of 10, less the one we hold in reserve. Each slot is taken as it's handed out, so consecutive callers
	// work their way through the burst rather than all being waved through.
	for i := 0; i < 9; i++ {
		require.Zero(t, p.reserveSlot(), "request %d is still within the burst", i+1)
	}

	require.Equal(t, emissionInterval, p.reserveSlot())
	require.Equal(t, twoEmissionIntervals, p.reserveSlot())
}

func TestPacer_AsTimePasses_TheAllowanceComesBack(t *testing.T) {
	p, now := newTestPacer()
	p.observeResponse(testPolicy, "r=0;t=10")
	require.Equal(t, twoEmissionIntervals, p.reserveSlot())

	// The slot we just took is ours; waiting it out (and one more interval) puts us back at the front of the queue.
	*now += 3 * emissionInterval

	require.Zero(t, p.reserveSlot())
}

func TestPacer_WhenAnOlderResponseArrivesLast_DoesNotTalkUsIntoSpeedingUp(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=1;t=9")

	// Responses can overtake each other, and this one is reporting a bucket that was fuller than the one we already
	// know about. Believing it would have us spend quota that has already gone.
	p.observeResponse(testPolicy, "r=9;t=1")

	require.Equal(t, emissionInterval, p.reserveSlot())
}

func TestPacer_WhenTheServerReportsLessQuotaThanWeExpected_SlowsUsDown(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=9;t=1")
	require.Zero(t, p.reserveSlot())

	// Something else has been spending the same allowance, so what's left isn't what our own arithmetic said.
	p.observeResponse(testPolicy, "r=0;t=10")

	require.Equal(t, twoEmissionIntervals, p.reserveSlot())
}

func TestPacer_WithABurstOfOne_LeavesNoRoomToBurstRatherThanGoingBackwards(t *testing.T) {
	p, _ := newTestPacer()

	// There is nothing to hold in reserve here, so the burst tolerance would come out negative if we let it.
	p.observeResponse("l=1;rpm=60", "r=1;t=0")
	require.Zero(t, p.reserveSlot())

	p.observeResponse("l=1;rpm=60", "r=0;t=1")
	require.Equal(t, emissionInterval, p.reserveSlot())
}

func TestPacer_WhenThePolicyChanges_StartsAgainOnTheNewOne(t *testing.T) {
	p, _ := newTestPacer()

	// Held back by a policy that only allows a request every 10 seconds...
	p.observeResponse("l=10;rpm=6", "r=1;t=90")
	require.Equal(t, 10*time.Second, p.reserveSlot())

	// ...which an administrator then loosens. The schedule we'd worked out is in the old policy's units, so holding on
	// to it would keep us slow long after the limit that justified it had gone.
	p.observeResponse(testPolicy, "r=10;t=0")

	require.Zero(t, p.reserveSlot())
}

func TestPacer_WithAPolicyItCannotMakeSenseOf_KeepsOutOfTheWay(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse("l=nonsense;rpm=0", "r=0;t=10")

	require.Zero(t, p.reserveSlot())
}

func TestPacer_WhenAResponseCarriesNoRemainingCount_GivesTheSlotBack(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=10;t=0")

	// Spend the burst, so the next request has to wait for the schedule we've built up.
	for i := 0; i < 10; i++ {
		p.reserveSlot()
	}

	require.Equal(t, twoEmissionIntervals, p.reserveSlot())

	// Two of those requests turn out to have gone somewhere the server doesn't rate limit, so they cost us nothing and
	// the schedule they claimed comes back. Otherwise we'd pace our whole request volume against a limit that only
	// governs part of it.
	p.observeResponse("", "")
	p.observeResponse(testPolicy, "")

	require.Equal(t, emissionInterval, p.reserveSlot())
}

func TestPacer_WhenRefundsOutnumberTheSlotsTaken_DoesNotBankCreditForAFutureBurst(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=10;t=0")
	p.reserveSlot()

	// Once the schedule is back level with the clock there's nothing left to give back, however many unlimited
	// responses arrive. Otherwise a quiet spell of unlimited requests would buy a burst we never earned.
	for i := 0; i < 20; i++ {
		p.observeResponse("", "")
	}

	// Still the burst the policy allows, and no longer.
	for i := 0; i < 9; i++ {
		require.Zero(t, p.reserveSlot(), "request %d is still within the burst", i+1)
	}

	require.Equal(t, emissionInterval, p.reserveSlot())
}

func TestPacer_AfterRefundingTooFar_PutsTheFloorBackUnderTheSchedule(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=0;t=10")

	// A request we refund may in fact have reached the limiter - a cancelled one, say. Refunding is the safe direction
	// to be wrong in, because the next response that does report a count resyncs us. Had these refunds stood, the
	// request below would have gone straight out.
	p.observeResponse("", "")
	p.observeResponse("", "")
	p.observeResponse("", "")
	p.observeResponse(testPolicy, "r=0;t=10")

	require.Equal(t, twoEmissionIntervals, p.reserveSlot())
}

func TestPacer_WhenTheServerStopsRateLimitingUs_StopsPacing(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=0;t=10")
	require.Equal(t, twoEmissionIntervals, p.reserveSlot())

	// An administrator turned the policy off mid-session, so responses stop carrying the headers. The refunds unwind
	// the schedule we'd built, rather than leaving us pacing for ever against a limit that has gone.
	for i := 0; i < 11; i++ {
		p.observeResponse("", "")
	}

	require.Zero(t, p.reserveSlot())
}

func TestGetParameter_ReadsTheValue(t *testing.T) {
	testCases := []struct {
		header   string
		name     string
		expected float64
	}{
		{"l=200;rpm=600", "l", 200},
		{"l=200;rpm=600", "rpm", 600},
		{"r=7;t=18", "t", 18},
		{"l=200; rpm=600", "rpm", 600},
		{"rpm=0.5", "rpm", 0.5},
		{"L=200", "l", 200},
		{"l=2e2", "l", 200},
	}

	for _, tc := range testCases {
		t.Run(tc.header+"/"+tc.name, func(t *testing.T) {
			value, ok := getParameter(tc.header, tc.name)
			require.True(t, ok)
			require.Equal(t, tc.expected, value)
		})
	}
}

func TestGetParameter_WhenItIsNotThere_SaysSo(t *testing.T) {
	testCases := []struct {
		header string
		name   string
	}{
		{"", "l"},
		{"l=200", "rpm"},
		{"l=200", ""},
		{"nonsense", "l"},
		{"l=", "l"},
		{"l=abc", "l"},
		// strconv.ParseFloat would read all of these; the server never sends them, so we don't either.
		{"l=Inf", "l"},
		{"l=NaN", "l"},
		{"l=0x10", "l"},
	}

	for _, tc := range testCases {
		t.Run(tc.header+"/"+tc.name, func(t *testing.T) {
			_, ok := getParameter(tc.header, tc.name)
			require.False(t, ok)
		})
	}
}

func TestGetParameter_IgnoresParametersItDoesNotKnowAbout(t *testing.T) {
	// So the server can add to these headers without older clients falling over.
	value, ok := getParameter("l=200;something=new;rpm=600", "rpm")
	require.True(t, ok)
	require.Equal(t, float64(600), value)
}
