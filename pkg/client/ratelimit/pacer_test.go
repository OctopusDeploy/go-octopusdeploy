package ratelimit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// A policy of 60 requests per minute makes one request cost exactly 1 second, which keeps the arithmetic in these
// tests easy to read. A burst limit of 10 gives 8 requests of burst room: one for the request that is sent, and one
// that the pacer keeps back.
const (
	testPolicy           = "l=10;rpm=60"
	testEmissionInterval = time.Second
	testBurstRoom        = 8
)

// newTestPacer returns a pacer with a clock that only moves when the test calls advance. Thus the arithmetic is exact.
func newTestPacer() (*pacer, func(time.Duration)) {
	var now time.Duration
	p := newPacerWithClock(func() time.Duration {
		return now
	})

	return p, func(duration time.Duration) {
		now += duration
	}
}

func TestPacerDoesNotPaceUntilTheServerTellsTheLimit(t *testing.T) {
	p, _ := newTestPacer()

	// An older server or a policy without a limit gets here, and neither must make the client slower.
	require.Equal(t, time.Duration(0), p.reserveSlot())
	require.Equal(t, time.Duration(0), p.reserveSlot())
}

func TestPacerDoesNotPaceWhenAResponseHasNoRateLimitHeaders(t *testing.T) {
	p, _ := newTestPacer()

	// A server that has no policy for this caller sends no headers, and the pacer must not make up a limit.
	p.observeResponse("", "")

	require.Equal(t, time.Duration(0), p.reserveSlot())
}

func TestPacerDoesNotWaitWhileRequestsRemain(t *testing.T) {
	testCases := []struct {
		name      string
		rateLimit string
	}{
		{"FullBucket", "r=10;t=0"},
		{"HalfBucket", "r=5;t=5"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			p, _ := newTestPacer()

			p.observeResponse(testPolicy, testCase.rateLimit)

			require.Equal(t, time.Duration(0), p.reserveSlot())
		})
	}
}

func TestPacerWaitsWhenOnlyTheReservedRequestRemains(t *testing.T) {
	p, _ := newTestPacer()

	// If the client spends the last request, the server has nothing left for the rounding difference between
	// its bucket and the pacer, thus the pacer waits for the bucket to refill one request first.
	p.observeResponse(testPolicy, "r=1;t=9")

	require.Equal(t, testEmissionInterval, p.reserveSlot())
}

func TestPacerWaitsForTheBucketToRefillWhenNoRequestsRemain(t *testing.T) {
	p, _ := newTestPacer()

	p.observeResponse(testPolicy, "r=0;t=10")

	// The pacer also waits for the request that it keeps back.
	require.Equal(t, 2*testEmissionInterval, p.reserveSlot())
}

func TestPacerSendsTheBurstThenPacesAtTheSteadyRate(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=10;t=0")

	// All of the requests in the burst go immediately, at the same time.
	for i := range testBurstRoom + 1 {
		require.Equal(t, time.Duration(0), p.reserveSlot(), "request %d is in the burst", i+1)
	}

	// After the burst, each request waits for the bucket to refill for it. The requests wait one emission
	// interval after each other, and they do not all get the same wait.
	require.Equal(t, testEmissionInterval, p.reserveSlot())
	require.Equal(t, 2*testEmissionInterval, p.reserveSlot())
	require.Equal(t, 3*testEmissionInterval, p.reserveSlot())
}

func TestPacerCountsTimeThatPassedAgainstTheWait(t *testing.T) {
	p, advance := newTestPacer()
	p.observeResponse(testPolicy, "r=0;t=10")

	require.Equal(t, 2*testEmissionInterval, p.reserveSlot())

	// The caller waited, thus the bucket refilled for the next request.
	advance(2 * testEmissionInterval)
	require.Equal(t, testEmissionInterval, p.reserveSlot())
	advance(3 * testEmissionInterval)
	require.Equal(t, time.Duration(0), p.reserveSlot())
}

func TestPacerSlowsDownWhenTheServerTellsFewerRequestsThanExpected(t *testing.T) {
	testCases := []struct {
		name          string
		rateLimit     string
		expectedDelay time.Duration
	}{
		{"OneRemains", "r=1;t=9", testEmissionInterval},
		{"NoneRemain", "r=0;t=10", 2 * testEmissionInterval},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			p, _ := newTestPacer()
			p.observeResponse(testPolicy, "r=10;t=0")
			require.Equal(t, time.Duration(0), p.reserveSlot())

			// Something else spends the same allowance, for example the web portal or a script of the same user.
			p.observeResponse(testPolicy, testCase.rateLimit)

			require.Equal(t, testCase.expectedDelay, p.reserveSlot())
		})
	}
}

func TestPacerDoesNotSpeedUpForAResponseThatArrivedOutOfOrder(t *testing.T) {
	testCases := []struct {
		name          string
		newer         string
		older         string
		expectedDelay time.Duration
	}{
		{"NoneRemain", "r=0;t=10", "r=10;t=0", 2 * testEmissionInterval},
		{"OneRemains", "r=1;t=9", "r=9;t=1", testEmissionInterval},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			p, _ := newTestPacer()
			p.observeResponse(testPolicy, testCase.newer)

			// A response that left the server before the newer one, but arrived after it, tells a bucket that is
			// fuller than it is. If the pacer used it, the client would spend requests that are already spent.
			p.observeResponse(testPolicy, testCase.older)

			require.Equal(t, testCase.expectedDelay, p.reserveSlot())
		})
	}
}

func TestPacerHandlesAPolicyWithNoRoomToBurst(t *testing.T) {
	p, _ := newTestPacer()

	// A burst limit of 1 gives a negative burst room after the request that is sent and the request that the
	// pacer keeps back. The burst room must be 0, so that the pacer does not pace into the past.
	p.observeResponse("l=1;rpm=60", "r=1;t=0")

	require.Equal(t, time.Duration(0), p.reserveSlot())
	require.Equal(t, testEmissionInterval, p.reserveSlot())
}

func TestPacerWaitsWithABurstOfOneWhenNoRequestsRemain(t *testing.T) {
	p, _ := newTestPacer()

	p.observeResponse("l=1;rpm=60", "r=1;t=0")
	require.Equal(t, time.Duration(0), p.reserveSlot())

	p.observeResponse("l=1;rpm=60", "r=0;t=1")
	require.Equal(t, testEmissionInterval, p.reserveSlot())
}

func TestPacerStartsAgainWhenThePolicyChanges(t *testing.T) {
	testCases := []struct {
		name          string
		oldPolicy     string
		oldRateLimit  string
		expectedDelay time.Duration
		newPolicy     string
		newRateLimit  string
	}{
		{"MoreRequests", testPolicy, "r=0;t=10", 2 * testEmissionInterval, "l=100;rpm=600", "r=100;t=0"},
		{"FasterRate", "l=10;rpm=6", "r=1;t=90", 10 * time.Second, testPolicy, "r=10;t=0"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			p, _ := newTestPacer()
			p.observeResponse(testCase.oldPolicy, testCase.oldRateLimit)
			require.Equal(t, testCase.expectedDelay, p.reserveSlot())

			// An administrator changes the policy. The old schedule uses the units of the old policy, thus it is
			// not correct now. The remaining requests in this response set the schedule again.
			p.observeResponse(testCase.newPolicy, testCase.newRateLimit)

			require.Equal(t, time.Duration(0), p.reserveSlot())
		})
	}
}

func TestPacerDoesNotPaceWhenThePolicyIsNotValid(t *testing.T) {
	p, _ := newTestPacer()

	p.observeResponse("l=nonsense;rpm=0", "r=0;t=10")

	// The pacer has no limit to pace against. It is better to send the request and let the limiter decide
	// than to make up a limit and keep back requests that the server would accept.
	require.Equal(t, time.Duration(0), p.reserveSlot())
}

func TestPacerKeepsThePolicyWhenALaterPolicyHeaderIsNotValid(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=10;t=0")

	p.observeResponse("", "r=0;t=10")

	require.Equal(t, 2*testEmissionInterval, p.reserveSlot())
}

func TestPacerGivesTheSlotBackWhenAResponseHasNoRemainingCount(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=10;t=0")

	// Spend the burst, so that the next request must wait for the schedule.
	for range testBurstRoom + 2 {
		p.reserveSlot()
	}
	require.Equal(t, 2*testEmissionInterval, p.reserveSlot())

	// Two of those requests went to endpoints that the server does not rate limit. They did not cost
	// anything, thus the pacer gives back the schedule that they took.
	p.observeResponse("", "")
	p.observeResponse(testPolicy, "")

	require.Equal(t, testEmissionInterval, p.reserveSlot())
}

func TestPacerDoesNotGiveCreditForAFutureBurstFromRefunds(t *testing.T) {
	testCases := []struct {
		name    string
		advance time.Duration
	}{
		// The schedule is level with the clock.
		{"ScheduleLevelWithTheClock", 0},
		// The schedule fell behind the clock.
		{"ScheduleBehindTheClock", 10 * testEmissionInterval},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			p, advance := newTestPacer()
			p.observeResponse(testPolicy, "r=10;t=0")
			p.reserveSlot()

			// When the schedule is not ahead of the clock, there is nothing to give back, however many unlimited
			// responses arrive. If not, a quiet period would give the client a burst that it did not earn.
			advance(testCase.advance)
			for range 20 {
				p.observeResponse("", "")
			}

			// The burst is the burst that the policy permits, and no longer.
			for i := range testBurstRoom + 1 {
				require.Equal(t, time.Duration(0), p.reserveSlot(), "request %d is in the burst", i+1)
			}
			require.Equal(t, testEmissionInterval, p.reserveSlot())
		})
	}
}

func TestPacerMovesTheScheduleBackAfterARefundThatIsTooLarge(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=0;t=10")

	// A request that the pacer refunds can be a request that got to the limiter, for example a cancelled request.
	// A refund that is too large is safe, because the next response with a remaining count sets the schedule
	// again. If these three refunds stayed, the request below would go immediately.
	p.observeResponse("", "")
	p.observeResponse("", "")
	p.observeResponse("", "")
	p.observeResponse(testPolicy, "r=0;t=10")

	require.Equal(t, 2*testEmissionInterval, p.reserveSlot())
}

func TestPacerStopsPacingWhenTheServerStopsTheRateLimit(t *testing.T) {
	p, _ := newTestPacer()
	p.observeResponse(testPolicy, "r=0;t=10")
	require.Equal(t, 2*testEmissionInterval, p.reserveSlot())

	// An administrator turns off the policy, thus the responses have no headers. The refunds remove the schedule,
	// so that the client does not pace against a limit that is gone.
	for range 11 {
		p.observeResponse("", "")
	}

	require.Equal(t, time.Duration(0), p.reserveSlot())
}

func TestPacerPacesAgainstAFractionalRequestRate(t *testing.T) {
	p, _ := newTestPacer()

	// Policies use requests per minute, thus a slow policy has a fraction.
	p.observeResponse("l=2;rpm=0.5", "r=0;t=240")

	require.Equal(t, 2*2*time.Minute, p.reserveSlot())
}

func TestGetParameterReadsTheValue(t *testing.T) {
	testCases := []struct {
		name     string
		header   string
		param    string
		expected float64
	}{
		{"BurstLimit", "l=200;rpm=600", "l", 200},
		{"RequestsPerMinute", "l=200;rpm=600", "rpm", 600},
		{"SecondsToRefill", "r=7;t=18", "t", 18},
		{"Fraction", "l=10;rpm=0.5", "rpm", 0.5},
		{"SpaceAfterSeparator", "l=200; rpm=600", "rpm", 600},
		{"SpacesAroundParameters", " l = 200 ; rpm = 600 ", "rpm", 600},
		{"NameIsNotCaseSensitive", "L=200", "l", 200},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			value, found := getParameter(testCase.header, testCase.param)

			require.True(t, found)
			require.Equal(t, testCase.expected, value)
		})
	}
}

func TestGetParameterReturnsFalseWhenTheValueIsNotThere(t *testing.T) {
	testCases := []struct {
		name   string
		header string
		param  string
	}{
		{"EmptyHeader", "", "l"},
		{"NoSuchParameter", "l=200;rpm=600", "r"},
		{"EmptyName", "l=200", ""},
		{"NoSeparator", "nonsense", "l"},
		{"NameOnly", "l", "l"},
		{"EmptyValue", "l=", "l"},
		{"NotANumber", "l=nonsense", "l"},
		// strconv.ParseFloat reads these, but they are not what the server means.
		{"Hexadecimal", "l=0x10", "l"},
		{"Infinity", "l=Inf", "l"},
		{"NotANumberLiteral", "l=NaN", "l"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, found := getParameter(testCase.header, testCase.param)

			require.False(t, found)
		})
	}
}

func TestGetParameterIgnoresParametersThatItDoesNotKnow(t *testing.T) {
	// So that the server can add parameters to these headers without a change to the client.
	value, found := getParameter("l=200;something=new;rpm=600", "rpm")

	require.True(t, found)
	require.Equal(t, float64(600), value)
}
