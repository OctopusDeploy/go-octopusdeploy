package examples

import (
	"fmt"
	"net/url"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/ratelimit"
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

	credentials, err := client.NewApiKey(apiKey)
	if err != nil {
		_ = fmt.Errorf("error creating API key credentials: %v", err)
		return
	}

	// start from the recommended settings: 3 retries, a maximum wait of 60 seconds,
	// and throttling with the advisory rate limit headers
	rateLimitOptions := ratelimit.DefaultOptions()

	// change the settings if necessary
	rateLimitOptions.MaxRetries = 5
	rateLimitOptions.MaxWait = 30 * time.Second

	options := &client.ClientOptions{
		RateLimit: &rateLimitOptions,
	}

	octopusClient, err := client.NewClientWithOptions(nil, apiURL, credentials, spaceID, "", options)
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
