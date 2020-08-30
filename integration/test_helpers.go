package integration

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	uuid "github.com/google/uuid"
)

var (
	octopusURL    string
	octopusAPIKey string
	client        *octopusdeploy.Client
)

func initTest() *octopusdeploy.Client {
	octopusURL = os.Getenv("OCTOPUS_URL")
	octopusAPIKey = os.Getenv("OCTOPUS_APIKEY")

	if octopusURL == "" || octopusAPIKey == "" {
		log.Fatal("Please make sure to set the env variables 'OCTOPUS_URL' and 'OCTOPUS_APIKEY' before running this test")
	}

	httpClient := http.Client{}
	client := octopusdeploy.NewClient(&httpClient, octopusURL, octopusAPIKey)

	return client
}

func getRandomName() string {
	fullName := fmt.Sprintf("go-octopusdeploy %s", uuid.New())
	fullName = fullName[0:49] //Some names in Octopus have a max limit of 50 characters (such as Environment Name)
	return fullName
}

func getRandomVarName() string {
	return fmt.Sprintf("go-octo-%v", time.Now().Unix())
}
