package octopusdeploy

type Form struct {
	Elements []*FormElement    `json:"Elements"`
	Values   map[string]string `json:"Values,omitempty"`
}
