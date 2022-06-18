package machines

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomDuration(mininum time.Duration) time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", rand.Int63n(1000)))
	duration += mininum
	return duration
}
