package machines

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomDuration(minimum int64) time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", rand.Int63n(1000)+minimum))
	return duration
}
