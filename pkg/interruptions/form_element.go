package interruptions

type FormElement struct {
	Control         Control `json:"Control,omitempty"`
	IsValueRequired *bool   `json:"IsValueRequired"`
	Name            string  `json:"Name,omitempty"`
}
