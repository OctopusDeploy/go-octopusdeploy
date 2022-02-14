package resources

type ReleaseProgression struct {
	Channel                 *Channel                   `json:"Channel,omitempty"`
	Deployments             map[string][]DashboardItem `json:"Deployments,omitempty"`
	HasUnresolvedDefect     bool                       `json:"HasUnresolvedDefect,omitempty"`
	NextDeployments         []string                   `json:"NextDeployments"`
	Release                 *Release                   `json:"Release,omitempty"`
	ReleaseRetentionPeriod  *RetentionPeriod           `json:"ReleaseRetentionPeriod,omitempty"`
	TentacleRetentionPeriod *RetentionPeriod           `json:"TentacleRetentionPeriod,omitempty"`
}
