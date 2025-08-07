package internal

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/dghubble/sling"
	"github.com/google/uuid"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetRandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		randomInt := rand.Intn(len(charset))
		result[i] = charset[randomInt]
	}
	return string(result)
}

func GetRandomName() string {
	fullName := fmt.Sprintf("test-id %s", uuid.New())
	fullName = fullName[0:44] //Some names in Octopus have a max limit of 50 characters (such as Environment Name)
	return fullName
}

func GetRandomThumbprint() string {
	thumbprint := strings.ToUpper(strings.ReplaceAll(fmt.Sprintf("%s%s", uuid.New(), uuid.New()), "-", ""))
	return thumbprint[0:40]
}

func GetRandomVersion() string {
	now := time.Now()
	return fmt.Sprintf("%d.%d.%d.%d", now.Year(), now.Month(), now.Day(), now.Hour()*10000+now.Minute()*100+now.Second())
}

func GetRandomPollingAddress() *url.URL {
	// Generate a random polling address using a random string
	randomString := strings.ToLower(GetRandomString(20))
	pollingAddress := fmt.Sprintf("poll://%s/", randomString)
	parsedURL, _ := url.Parse(pollingAddress)
	return parsedURL
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
