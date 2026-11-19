package machinepolicies

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ToTimeSpan(duration time.Duration) string {
	days := duration / (time.Minute * 1440)
	duration -= days * (time.Minute * 1440)
	hours := duration / time.Hour
	duration -= hours * time.Hour
	minutes := duration / time.Minute
	duration -= minutes * time.Minute
	seconds := duration / time.Second
	duration -= seconds * time.Second
	secondsFraction := duration.Milliseconds() * 100

	if secondsFraction == 0 {
		if days == 0 {
			return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
		}
		return fmt.Sprintf("%02d.%02d:%02d:%02d", days, hours, minutes, seconds)
	}

	if days == 0 {
		return fmt.Sprintf("%02d:%02d:%02d.%05d", hours, minutes, seconds, secondsFraction)
	}
	return fmt.Sprintf("%02d.%02d:%02d:%02d.%05d", days, hours, minutes, seconds, secondsFraction)
}

// FromTimeSpan parses a .NET TimeSpan ("[-][d.]hh:mm:ss[.fffffff]") into a duration.
func FromTimeSpan(timeSpan string) time.Duration {
	if timeSpan == "" {
		return 0
	}

	negative := strings.HasPrefix(timeSpan, "-")
	if negative {
		timeSpan = timeSpan[1:]
	}

	var days int64
	if dot := strings.IndexByte(timeSpan, '.'); dot >= 0 && dot < strings.IndexByte(timeSpan, ':') {
		days, _ = strconv.ParseInt(timeSpan[:dot], 10, 64)
		timeSpan = timeSpan[dot+1:]
	}

	var fraction time.Duration
	if dot := strings.IndexByte(timeSpan, '.'); dot >= 0 {
		// .NET fractional seconds are ticks of 100ns, up to 7 digits.
		ticks := timeSpan[dot+1:]
		if len(ticks) > 7 {
			ticks = ticks[:7]
		}
		ticks += strings.Repeat("0", 7-len(ticks))
		parsed, _ := strconv.ParseInt(ticks, 10, 64)
		fraction = time.Duration(parsed) * 100 * time.Nanosecond
		timeSpan = timeSpan[:dot]
	}

	parts := strings.Split(timeSpan, ":")
	if len(parts) != 3 {
		return 0
	}
	hours, _ := strconv.ParseInt(parts[0], 10, 64)
	minutes, _ := strconv.ParseInt(parts[1], 10, 64)
	seconds, _ := strconv.ParseInt(parts[2], 10, 64)

	duration := time.Duration(days)*24*time.Hour +
		time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second +
		fraction

	if negative {
		return -duration
	}
	return duration
}
