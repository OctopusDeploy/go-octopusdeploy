package interruptions

type InterruptionSubmitRequest struct {
	Instructions string `json:"Instructions"`
	Notes        string `json:"Notes"`
	Result       string `json:"Result"`
}

func NewInterruptionSubmitRequest() *InterruptionSubmitRequest {
	return &InterruptionSubmitRequest{}
}
