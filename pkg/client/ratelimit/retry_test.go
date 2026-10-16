package ratelimit

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// noJitter makes retryDelay add no jitter, so that the arithmetic is easy to read. Jitter has its own tests.
const noJitter = 0

// maxJitter makes retryDelay add almost all of the jitter that it can add.
const maxJitter = 0.999

func retryOptions() Options {
	return Options{
		RetryCount:   5,
		DefaultDelay: 2 * time.Second,
		MaxDelay:     30 * time.Second,
	}
}

func TestRetryDelayWaitsForTheSecondsInRetryAfter(t *testing.T) {
	delay, ok := retryDelay("5", retryOptions(), 0, noJitter)

	require.True(t, ok)
	require.Equal(t, 5*time.Second, delay)
}

func TestRetryDelayWaitsTheDefaultDelay(t *testing.T) {
	testCases := []struct {
		name       string
		retryAfter string
	}{
		{"MissingRetryAfter", ""},
		{"RetryAfterIsNotANumber", "soon please"},
		// A malformed delay waits the default delay, and does not retry immediately.
		{"Fraction", "5.5"},
		{"PlusSign", "+5"},
		{"FractionWithoutWholeNumber", ".5"},
		// retryDelay does not read Retry-After as a date. A client clock that is wrong by some minutes would
		// give a delay of 0 and send too many requests, or a delay longer than MaxDelay and stop the retries.
		{"FutureDate", "Sun, 30 Aug 2026 10:00:07 GMT"},
		{"PastDate", "Sun, 30 Aug 2026 09:59:50 GMT"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			delay, ok := retryDelay(testCase.retryAfter, retryOptions(), 0, noJitter)

			require.True(t, ok)
			require.Equal(t, 2*time.Second, delay)
		})
	}
}

func TestRetryDelayRetriesImmediatelyWhenRetryAfterIsZeroOrNegative(t *testing.T) {
	testCases := []struct {
		name       string
		retryAfter string
	}{
		{"Zero", "0"},
		{"Negative", "-3"},
		{"VeryNegative", "-99999999999999999999"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			delay, ok := retryDelay(testCase.retryAfter, retryOptions(), 0, noJitter)

			require.True(t, ok)
			require.Equal(t, time.Duration(0), delay)
		})
	}
}

func TestRetryDelayGivesUpWhenTheServerTellsTheClientToWaitLongerThanMaxDelay(t *testing.T) {
	_, ok := retryDelay("31", retryOptions(), 0, noJitter)
	require.False(t, ok)

	delay, ok := retryDelay("30", retryOptions(), 0, noJitter)
	require.True(t, ok)
	require.Equal(t, 30*time.Second, delay)
}

func TestRetryDelayGivesUpWhenRetryAfterIsTooLargeForADuration(t *testing.T) {
	testCases := []struct {
		name       string
		retryAfter string
	}{
		{"LargerThanADuration", "9999999999999"},
		{"LargerThanAnInt64", "99999999999999999999"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, ok := retryDelay(testCase.retryAfter, retryOptions(), 0, noJitter)

			require.False(t, ok)
		})
	}
}

func TestRetryDelayDoublesTheWaitForEachRetry(t *testing.T) {
	// The server tells all of the requests in a burst that it rejects to wait the same time. A retry after that
	// time only lets through the requests that the limiter refills in that time, thus the client backs off to
	// give the limiter time to refill for all of the requests that it rejected.
	expectedDelays := []time.Duration{2 * time.Second, 4 * time.Second, 8 * time.Second, 16 * time.Second}

	for attempt, expectedDelay := range expectedDelays {
		delay, ok := retryDelay("2", retryOptions(), attempt, noJitter)

		require.True(t, ok)
		require.Equal(t, expectedDelay, delay, "attempt %d", attempt)
	}
}

func TestRetryDelayDoesNotBackOffPastMaxDelay(t *testing.T) {
	delay, ok := retryDelay("2", retryOptions(), 10, noJitter)

	require.True(t, ok)
	require.Equal(t, 30*time.Second, delay)
}

func TestRetryDelayDoesNotOverflowForLargeAttempts(t *testing.T) {
	delay, ok := retryDelay("2", retryOptions(), 1000, noJitter)

	require.True(t, ok)
	require.Equal(t, 30*time.Second, delay)
}

func TestRetryDelayBacksOffFromTheDefaultDelayWhenRetryAfterIsMissing(t *testing.T) {
	delay, ok := retryDelay("", retryOptions(), 2, noJitter)

	require.True(t, ok)
	require.Equal(t, 8*time.Second, delay)
}

func TestRetryDelayAddsJitterAfterTheWaitNeverBeforeIt(t *testing.T) {
	// The server told the earliest time to retry, thus jitter only makes the delay longer.
	delay, ok := retryDelay("2", retryOptions(), 0, maxJitter)

	require.True(t, ok)
	require.Greater(t, delay, 2*time.Second)
	require.LessOrEqual(t, delay, 2500*time.Millisecond)
}

func TestRetryDelayKeepsTheJitterInMaxDelay(t *testing.T) {
	delay, ok := retryDelay("2", retryOptions(), 10, maxJitter)
	require.True(t, ok)
	require.Equal(t, 30*time.Second, delay)

	delay, ok = retryDelay("28", retryOptions(), 0, maxJitter)
	require.True(t, ok)
	require.LessOrEqual(t, delay, 30*time.Second)
}

func TestWaitReturnsWhenTheDelayIsOver(t *testing.T) {
	require.NoError(t, wait(context.Background(), 0))
}

func TestWaitReturnsTheErrorWhenTheContextIsAlreadyCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	require.ErrorIs(t, wait(ctx, 0), context.Canceled)
}

func TestWaitReturnsTheErrorWhenTheContextIsCancelledDuringTheWait(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(10*time.Millisecond, cancel)

	require.ErrorIs(t, wait(ctx, time.Minute), context.Canceled)
}
