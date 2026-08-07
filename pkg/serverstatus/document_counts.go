package serverstatus

type DocumentCounts struct {
	Global         GlobalDocumentCounts         `json:"Global"`
	Infrastructure InfrastructureDocumentCounts `json:"Infrastructure"`
	Library        LibraryDocumentCounts        `json:"Library"`
	Project        ProjectDocumentCounts        `json:"Project"`
}

type GlobalDocumentCounts struct {
	Spaces int `json:"Spaces"`
	Teams  int `json:"Teams"`
	Users  int `json:"Users"`
}

type InfrastructureDocumentCounts struct {
	DeploymentTargets int `json:"DeploymentTargets"`
	Environments      int `json:"Environments"`
	Tenants           int `json:"Tenants"`
	WorkerPools       int `json:"WorkerPools"`
	Workers           int `json:"Workers"`
}

type LibraryDocumentCounts struct {
	Certificates int `json:"Certificates"`
	Packages     int `json:"Packages"`
	VariableSets int `json:"VariableSets"`
}

type ProjectDocumentCounts struct {
	Deployments int `json:"Deployments"`
	Projects    int `json:"Projects"`
	Releases    int `json:"Releases"`
	RunbookRuns int `json:"RunbookRuns"`
	Runbooks    int `json:"Runbooks"`
}
