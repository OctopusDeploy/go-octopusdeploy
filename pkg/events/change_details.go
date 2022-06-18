package events

type ChangeDetails struct {
	Differences     interface{} `json:"Differences,omitempty"`
	DocumentContext interface{} `json:"DocumentContext,omitempty"`
}
