package octopusdeploy

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
	t.Logf("500ms: %s", ToTimeSpan(halfSecond))
	t.Logf("1000ms: %s", ToTimeSpan(second))
	t.Logf("1s: %s", ToTimeSpan(time.Second))
	t.Logf("1m: %s", ToTimeSpan(time.Minute))
	t.Logf("1h: %s", ToTimeSpan(time.Hour))
	t.Logf("120m: %s", ToTimeSpan(twoHours))
	t.Logf("47h: %s", ToTimeSpan(fourtySevenHours))
	t.Logf("48h: %s", ToTimeSpan(twoDays))
}

func TestFromTimeSpan(t *testing.T) {
	t.Logf("1s: %s", FromTimeSpan("00.00:00:01"))
	t.Logf("1m: %s", FromTimeSpan("00.00:01:00"))
	t.Logf("1h: %s", FromTimeSpan("00.01:00:00"))
	t.Logf("120m: %s", FromTimeSpan("00.02:00:00"))
	t.Logf("47h: %s", FromTimeSpan("00.47:00:00"))
	t.Logf("48h: %s", FromTimeSpan("00.48:00:00"))
	t.Logf("2d: %s", FromTimeSpan("02.00:00:00"))
}
