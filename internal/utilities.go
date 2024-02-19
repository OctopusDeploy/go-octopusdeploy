package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/dghubble/sling"
	"github.com/google/uuid"
)

func GetRandomName() string {
	fullName := fmt.Sprintf("test-id %s", uuid.New())
	fullName = fullName[0:44] //Some names in Octopus have a max limit of 50 characters (such as Environment Name)
	return fullName
}

func GetRandomThumbprint() string {
	thumbprint := strings.ToUpper(strings.ReplaceAll(fmt.Sprintf("%s%s", uuid.New(), uuid.New()), "-", ""))
	return thumbprint
}

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func GetDefaultClient() *sling.Sling {
	host := os.Getenv(constants.EnvironmentVariableOctopusHost)
	apiKey := os.Getenv(constants.EnvironmentVariableOctopusApiKey)

	if len(host) == 0 {
		host = os.Getenv(constants.ClientURLEnvironmentVariable)
	}

	if len(apiKey) == 0 {
		apiKey = os.Getenv(constants.ClientAPIKeyEnvironmentVariable)
	}

	// NOTE: You can direct traffic through a proxy trace like Fiddler
	// Everywhere by preconfiguring the client to route traffic through a
	// proxy.

	// proxyStr := "http://127.0.0.1:5555"
	// proxyURL, _ := url.Parse(proxyStr)

	// tr := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyURL),
	// }
	// httpClient := http.Client{Transport: tr}

	return sling.New().Client(nil).Base(host).Set(constants.ClientAPIKeyHTTPHeader, apiKey)
}

func GetSpaceID(spaceID string, clientSpaceID string) (string, error) {
	if spaceID != "" {
		return spaceID, nil
	}

	if clientSpaceID != "" {
		return clientSpaceID, nil
	}

	return "", MissingSpaceIDError()
}

func TrimTemplate(uri string) string {
	return strings.TrimRight(strings.Split(uri, "{")[0], "/")
}
