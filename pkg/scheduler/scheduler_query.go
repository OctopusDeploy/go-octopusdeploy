package scheduler

type SchedulerQuery struct {
	Verbose bool   `uri:"verbose,omitempty" url:"verbose,omitempty"`
	Tail    string `uri:"tail,omitempty" url:"tail,omitempty"`
}
