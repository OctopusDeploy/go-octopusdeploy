package model

type ConnectivityCheck struct {
	DependsOnPropertyNames []string `json:"DependsOnPropertyNames"`
	Title                  string   `json:"Title,omitempty"`
	URL                    string   `json:"Url,omitempty"`
}
