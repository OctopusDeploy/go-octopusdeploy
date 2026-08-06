package machines

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

// FromTimeSpan parses a .NET time span, "[d.]hh:mm:ss[.fffffff]", into a duration. The day and
// fractional-second components are optional, and the server does not pad the day component to a
// fixed width, so the fields cannot be read at fixed offsets. Input that does not parse yields a
// zero duration.
func FromTimeSpan(timeSpan string) time.Duration {
	var days int64
	remainder := timeSpan

	// Both the day separator and the fractional-second separator are ".", so a leading segment is
	// only the day component when the rest still holds a complete "hh:mm:ss".
	if index := strings.Index(remainder, "."); index >= 0 && strings.Count(remainder[index+1:], ":") == 2 {
		parsedDays, err := strconv.ParseInt(remainder[:index], 10, 64)
		if err != nil {
			return 0
		}
		days = parsedDays
		remainder = remainder[index+1:]
	}

	var fraction time.Duration
	if index := strings.Index(remainder, "."); index >= 0 {
		digits := remainder[index+1:]
		parsedFraction, err := strconv.ParseInt(digits, 10, 64)
		if err != nil {
			return 0
		}
		// The digits are a decimal fraction of a second, however many of them there are.
		scale := pow10(len(digits))
		fraction = time.Duration(parsedFraction * int64(time.Second) / scale)
		remainder = remainder[:index]
	}

	fields := strings.Split(remainder, ":")
	if len(fields) != 3 {
		return 0
	}

	hours, err := strconv.ParseInt(fields[0], 10, 64)
	if err != nil {
		return 0
	}
	minutes, err := strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return 0
	}
	seconds, err := strconv.ParseInt(fields[2], 10, 64)
	if err != nil {
		return 0
	}

	return time.Duration(days)*24*time.Hour +
		time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second +
		fraction
}

func pow10(exponent int) int64 {
	result := int64(1)
	for i := 0; i < exponent; i++ {
		result *= 10
	}
	return result
}
