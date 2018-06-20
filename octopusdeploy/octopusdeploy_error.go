package octopusdeploy

import "fmt"

type OctopusDeployError struct {
	ErrorMessage  string   `json:"ErrorMessage"`
	Errors        []string `json:"Errors"`
	FullException string   `json:"FullException"`
}

func (e OctopusDeployError) Error() string {
	return fmt.Sprintf("Octopus Deploy Error Response: %v %+v %v", e.ErrorMessage, e.Errors, e.FullException)
}
