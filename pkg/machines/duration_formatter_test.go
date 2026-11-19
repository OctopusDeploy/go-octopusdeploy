package machines

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFromTimeSpanParsesDayComponent(t *testing.T) {
	assert.Equal(t, 24*time.Hour, FromTimeSpan("1.00:00:00"))
	assert.Equal(t, 24*time.Hour, FromTimeSpan("01.00:00:00"))
	assert.Equal(t, 48*time.Hour, FromTimeSpan("2.00:00:00"))
	assert.Equal(t, 25*time.Hour+30*time.Minute, FromTimeSpan("1.01:30:00"))
	// The server does not pad the day component, so it can be arbitrarily wide.
	assert.Equal(t, 900000*time.Hour, FromTimeSpan("37500.00:00:00"))
}

func TestFromTimeSpanParsesTimeOnly(t *testing.T) {
	assert.Equal(t, time.Duration(0), FromTimeSpan("00:00:00"))
	assert.Equal(t, time.Second, FromTimeSpan("00:00:01"))
	assert.Equal(t, 12*time.Hour+30*time.Minute+15*time.Second, FromTimeSpan("12:30:15"))
	assert.Equal(t, 47*time.Hour, FromTimeSpan("47:00:00"))
}

func TestFromTimeSpanParsesFractionalSeconds(t *testing.T) {
	assert.Equal(t, 500*time.Millisecond, FromTimeSpan("00:00:00.5000000"))
	// ToTimeSpan emits a 5-digit fraction rather than .NET's 7, so it must round-trip too.
	assert.Equal(t, 500*time.Millisecond, FromTimeSpan("00:00:00.50000"))
	assert.Equal(t, time.Second+500*time.Millisecond, FromTimeSpan("1.00:00:01.50000")-24*time.Hour)
	assert.Equal(t, 24*time.Hour+500*time.Millisecond, FromTimeSpan("01.00:00:00.50000"))
}

func TestFromTimeSpanHandlesEmptyAndMalformed(t *testing.T) {
	assert.Equal(t, time.Duration(0), FromTimeSpan(""))
	assert.Equal(t, time.Duration(0), FromTimeSpan("not-a-timespan"))
	// Too few and too many colon-separated fields.
	assert.Equal(t, time.Duration(0), FromTimeSpan("00:00"))
	assert.Equal(t, time.Duration(0), FromTimeSpan("00:00:00:00"))
}

func TestFromTimeSpanHandlesNegative(t *testing.T) {
	assert.Equal(t, -90*time.Minute, FromTimeSpan("-01:30:00"))
	assert.Equal(t, -25*time.Hour, FromTimeSpan("-1.01:00:00"))
}

func TestToTimeSpanRendersDotNetFormat(t *testing.T) {
	// The day and fractional-second components are only rendered when non-zero, and the
	// round-trip test cannot tell a formatting change from a matching parsing change.
	for _, c := range []struct {
		duration time.Duration
		expected string
	}{
		{500 * time.Millisecond, "00:00:00.50000"},
		{1111 * time.Millisecond, "00:00:01.11100"},
		{time.Second, "00:00:01"},
		{time.Minute, "00:01:00"},
		{time.Hour, "01:00:00"},
		{2 * time.Hour, "02:00:00"},
		{47 * time.Hour, "01.23:00:00"},
		{48 * time.Hour, "02.00:00:00"},
	} {
		assert.Equal(t, c.expected, ToTimeSpan(c.duration), "ToTimeSpan(%s)", c.duration)
	}
}

func TestToTimeSpanFromTimeSpanRoundTrip(t *testing.T) {
	for _, d := range []time.Duration{
		0,
		time.Second,
		time.Minute,
		time.Hour,
		24 * time.Hour,
		47 * time.Hour,
		48 * time.Hour,
		25*time.Hour + 30*time.Minute + 15*time.Second,
	} {
		assert.Equal(t, d, FromTimeSpan(ToTimeSpan(d)), "round trip failed for %s (rendered %q)", d, ToTimeSpan(d))
	}
}
