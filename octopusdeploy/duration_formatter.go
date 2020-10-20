package octopusdeploy

import (
	"fmt"
	"strconv"
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

func FromTimeSpan(timeSpan string) time.Duration {
	if len(timeSpan) == 8 {
		hours, _ := strconv.ParseInt(timeSpan[0:2], 10, 64)
		minutes, _ := strconv.ParseInt(timeSpan[3:5], 10, 64)
		seconds, _ := strconv.ParseInt(timeSpan[6:8], 10, 64)
		duration, _ := time.ParseDuration(fmt.Sprintf("%dh%dm%ds", hours, minutes, seconds))
		return duration
	}

	days, _ := strconv.ParseInt(timeSpan[0:0], 10, 32)
	hours, _ := strconv.ParseInt(timeSpan[2:4], 10, 64)
	hours += (days * 24)
	minutes, _ := strconv.ParseInt(timeSpan[5:7], 10, 64)
	seconds, _ := strconv.ParseInt(timeSpan[8:10], 10, 64)
	duration, _ := time.ParseDuration(fmt.Sprintf("%dh%dm%ds", hours, minutes, seconds))
	return duration
}
