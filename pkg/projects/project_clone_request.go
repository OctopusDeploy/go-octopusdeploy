package projects

type ProjectCloneRequest struct {
	Name           string `json:"Name"`
	Description    string `json:"Description"`
	ProjectGroupID string `json:"ProjectGroupID"`
	LifecycleID    string `json:"LifecycleID"`
}

type ProjectCloneQuery struct {
	CloneProjectID string `uri:"clone" url:"clone"`
}
