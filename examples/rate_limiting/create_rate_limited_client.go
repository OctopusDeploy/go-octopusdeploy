package examples

import (
	"fmt"
	"net/url"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client/ratelimit"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/spaces"
)

// CreateRateLimitedClientExample provides an example of how to create a Go API client
// that obeys the rate limits of Octopus Server.
func CreateRateLimitedClientExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	// start from the recommended settings: 5 retries, a default delay of 2 seconds,
	// a maximum delay of 30 seconds, and pacing with the rate limit headers
	rateLimitOptions := ratelimit.DefaultOptions()

	// change the settings if necessary
	rateLimitOptions.RetryCount = 10
	rateLimitOptions.MaxDelay = 60 * time.Second

	// the first parameter for ratelimit.NewHttpClient can accept a http.Client if you wish to
	// override the default; ratelimit.NewHttpClient returns a copy and does not change it
	httpClient := ratelimit.NewHttpClient(nil, rateLimitOptions)

	octopusClient, err := client.NewClient(httpClient, apiURL, apiKey, spaceID)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// all requests from this client now obey the rate limits
	space, err := spaces.GetByID(octopusClient, spaceID)
	if err != nil {
		_ = fmt.Errorf("error getting space: %v", err)
		return
	}

	fmt.Printf("space: (%s)\n", space.GetID())
}
