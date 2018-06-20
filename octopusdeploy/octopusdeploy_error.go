package octopusdeploy

import "fmt"

type APIError struct {
	ErrorMessage  string   `json:"ErrorMessage"`
	Errors        []string `json:"Errors"`
	FullException string   `json:"FullException"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Octopus Deploy Error Response: %v %+v %v", e.ErrorMessage, e.Errors, e.FullException)
}
