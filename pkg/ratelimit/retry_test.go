package ratelimit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Jitter is exercised on its own below; everywhere else we pin it out so the arithmetic is readable.
const noJitter = 0

func requireRetryDelay(t *testing.T, expected time.Duration, retryAfter string, options Options, attempt int) {
	t.Helper()
	delay, ok := retryDelay(retryAfter, options, attempt, noJitter)
	require.True(t, ok, "expected to retry after Retry-After %q", retryAfter)
	require.Equal(t, expected, delay)
}

func TestRetryDelay_WaitsForTheNumberOfSecondsInRetryAfter(t *testing.T) {
	requireRetryDelay(t, 5*time.Second, "5", Options{}, 0)
	requireRetryDelay(t, 5*time.Second, " 5 ", Options{}, 0)
}

func TestRetryDelay_WaitsTheDefaultDelayWhenTheServerDidNotSendARetryAfter(t *testing.T) {
	requireRetryDelay(t, 2*time.Second, "", Options{}, 0)
	requireRetryDelay(t, 3*time.Second, "", Options{RetryDefaultDelay: 3 * time.Second}, 0)
}

func TestRetryDelay_WaitsTheDefaultDelayWhenRetryAfterMakesNoSense(t *testing.T) {
	requireRetryDelay(t, 2*time.Second, "soon please", Options{}, 0)
}

func TestRetryDelay_WaitsTheDefaultDelayForAMalformedDelayRatherThanRetryingImmediately(t *testing.T) {
	requireRetryDelay(t, 2*time.Second, "5.5", Options{}, 0)
	requireRetryDelay(t, 2*time.Second, "+5", Options{}, 0)
	requireRetryDelay(t, 2*time.Second, ".5", Options{}, 0)
}

func TestRetryDelay_RetriesImmediatelyWhenRetryAfterIsAZeroOrNegativeDelay(t *testing.T) {
	requireRetryDelay(t, 0, "0", Options{}, 0)
	requireRetryDelay(t, 0, "-3", Options{}, 0)
}

func TestRetryDelay_WaitsTheDefaultDelayForADateFormRetryAfterRatherThanTrustingTheLocalClock(t *testing.T) {
	// We don't read Retry-After as a date. A machine running a few minutes out either way would compute a delay of zero
	// and hammer the server, or one past the max delay that silently stops us retrying.
	requireRetryDelay(t, 2*time.Second, "Sun, 30 Aug 2026 10:00:07 GMT", Options{}, 0)
	requireRetryDelay(t, 2*time.Second, "Sun, 30 Aug 2026 09:59:50 GMT", Options{}, 0)
}

func TestRetryDelay_GivesUpWhenTheServerAsksUsToWaitLongerThanWeAreWillingTo(t *testing.T) {
	_, ok := retryDelay("31", Options{}, 0, noJitter)
	require.False(t, ok)

	requireRetryDelay(t, 30*time.Second, "30", Options{}, 0)

	_, ok = retryDelay("11", Options{RetryMaxDelay: 10 * time.Second}, 0, noJitter)
	require.False(t, ok)
}

func TestRetryDelay_GivesUpWhenRetryAfterIsTooBigToRepresent(t *testing.T) {
	// Rather than overflowing into a short (or negative) wait.
	_, ok := retryDelay("9223372036854775807", Options{}, 0, noJitter)
	require.False(t, ok)

	_, ok = retryDelay("99999999999999999999999", Options{}, 0, noJitter)
	require.False(t, ok)
}

func TestRetryDelay_DoublesTheWaitForEachRetryWeHaveAlreadyMade(t *testing.T) {
	// A burst that is rate limited all at once is told to wait the same number of seconds by every rejection. Retrying
	// on that interval only lets through what the limiter refills in one interval, so we back off to give it time to
	// refill for everything it rejected.
	requireRetryDelay(t, 2*time.Second, "2", Options{}, 0)
	requireRetryDelay(t, 4*time.Second, "2", Options{}, 1)
	requireRetryDelay(t, 8*time.Second, "2", Options{}, 2)
	requireRetryDelay(t, 16*time.Second, "2", Options{}, 3)
}

func TestRetryDelay_NeverBacksOffPastTheLongestWaitWeAreWillingToMake(t *testing.T) {
	requireRetryDelay(t, 30*time.Second, "2", Options{}, 10)
	requireRetryDelay(t, 30*time.Second, "2", Options{}, 1000)
}

func TestRetryDelay_BacksOffFromTheDefaultDelayWhenTheServerDidNotSendARetryAfter(t *testing.T) {
	requireRetryDelay(t, 8*time.Second, "", Options{}, 2)
}

func TestRetryDelay_StillRetriesImmediatelyAtEveryAttemptWhenRetryAfterIsZero(t *testing.T) {
	requireRetryDelay(t, 0, "0", Options{}, 3)
}

func TestRetryDelay_AddsJitterAfterTheWaitNeverBeforeIt(t *testing.T) {
	// The server told us the earliest it wants us back, so jitter only ever makes us later.
	delay, ok := retryDelay("2", Options{}, 0, 0.999)
	require.True(t, ok)
	require.Greater(t, delay, 2*time.Second)
	require.LessOrEqual(t, delay, 2500*time.Millisecond)
}

func TestRetryDelay_KeepsTheJitterInsideTheLongestWaitWeAreWillingToMake(t *testing.T) {
	delay, ok := retryDelay("2", Options{}, 10, 0.999)
	require.True(t, ok)
	require.Equal(t, 30*time.Second, delay)

	delay, ok = retryDelay("28", Options{}, 0, 0.999)
	require.True(t, ok)
	require.LessOrEqual(t, delay, 30*time.Second)
}

func TestOptions_ZeroValueGivesTheDefaults(t *testing.T) {
	// These are what ships: five retries, two seconds when the server didn't say, and never more than thirty. The same
	// as the web portal.
	options := Options{}
	require.Equal(t, 5, options.retryCount())
	require.Equal(t, 2*time.Second, options.retryDefaultDelay())
	require.Equal(t, 30*time.Second, options.retryMaxDelay())
	require.False(t, options.UseRateLimitHeaders)
}

func TestOptions_ANegativeRetryCountTurnsRetryingOff(t *testing.T) {
	require.Equal(t, 0, Options{RetryCount: -1}.retryCount())
	require.Equal(t, 2, Options{RetryCount: 2}.retryCount())
}
