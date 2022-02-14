package resources

type EventReference struct {
	Length               int32  `json:"Length,omitempty"`
	ReferencedDocumentID string `json:"ReferencedDocumentId,omitempty"`
	StartIndex           int32  `json:"StartIndex,omitempty"`
}
