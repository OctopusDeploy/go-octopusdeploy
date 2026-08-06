package machinepolicies

import (
	"testing"
	"time"
)

func TestToTimeSpan(t *testing.T) {
	halfSecond, _ := time.ParseDuration("0.5s")
	second, _ := time.ParseDuration("1111ms")
	twoHours, _ := time.ParseDuration("120m")
	fourtySevenHours, _ := time.ParseDuration("47h")
	twoDays, _ := time.ParseDuration("48h")

	testCases := []struct {
		duration time.Duration
		expected string
	}{
		{halfSecond, "00:00:00.50000"},
		{second, "00:00:01.11100"},
		{time.Second, "00:00:01"},
		{time.Minute, "00:01:00"},
		{time.Hour, "01:00:00"},
		{twoHours, "02:00:00"},
		{fourtySevenHours, "01.23:00:00"},
		{twoDays, "02.00:00:00"},
	}

	for _, testCase := range testCases {
		if actual := ToTimeSpan(testCase.duration); actual != testCase.expected {
			t.Errorf("ToTimeSpan(%s) = %q, want %q", testCase.duration, actual, testCase.expected)
		}
	}
}

func TestFromTimeSpan(t *testing.T) {
	testCases := []struct {
		timeSpan string
		expected time.Duration
	}{
		// hh:mm:ss
		{"00:00:00", 0},
		{"00:00:01", time.Second},
		{"00:01:00", time.Minute},
		{"01:00:00", time.Hour},
		{"02:00:00", 2 * time.Hour},
		// d.hh:mm:ss, as written by ToTimeSpan
		{"00.00:00:01", time.Second},
		{"00.47:00:00", 47 * time.Hour},
		{"01.23:00:00", 47 * time.Hour},
		{"02.00:00:00", 48 * time.Hour},
		// d.hh:mm:ss, as returned by the Octopus server, which does not pad the days
		{"1.00:00:00", 24 * time.Hour},
		{"7.12:30:00", 7*24*time.Hour + 12*time.Hour + 30*time.Minute},
		{"37500.00:00:00", 900000 * time.Hour},
		// fractional seconds
		{"00:00:00.50000", 500 * time.Millisecond},
		{"00:00:01.11100", 1111 * time.Millisecond},
		{"01.00:00:00.50000", 24*time.Hour + 500*time.Millisecond},
		// .NET renders fractional seconds as seven digits
		{"00:00:00.5000000", 500 * time.Millisecond},
		// malformed input yields a zero duration rather than a panic
		{"", 0},
		{"not-a-timespan", 0},
		{"00:00", 0},
	}

	for _, testCase := range testCases {
		if actual := FromTimeSpan(testCase.timeSpan); actual != testCase.expected {
			t.Errorf("FromTimeSpan(%q) = %s, want %s", testCase.timeSpan, actual, testCase.expected)
		}
	}
}

func TestTimeSpanRoundTrip(t *testing.T) {
	durations := []time.Duration{
		0,
		time.Second,
		time.Minute,
		time.Hour,
		47 * time.Hour,
		48 * time.Hour,
		7 * 24 * time.Hour,
		900000 * time.Hour,
		500 * time.Millisecond,
		24*time.Hour + 12*time.Hour + 30*time.Minute + 15*time.Second,
	}

	for _, duration := range durations {
		if actual := FromTimeSpan(ToTimeSpan(duration)); actual != duration {
			t.Errorf("FromTimeSpan(ToTimeSpan(%s)) = %s, want %s", duration, actual, duration)
		}
	}
}
