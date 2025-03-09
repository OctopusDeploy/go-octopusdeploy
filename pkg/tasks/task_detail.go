package tasks

import "time"

type TaskDetailsResource struct {
	Task            *Task              `json:"Task"`
	ActivityLogs    []*ActivityElement `json:"ActivityLogs"`
	PhysicalLogSize int64              `json:"PhysicalLogSize"`
	Progress        *TaskProgress      `json:"Progress"`
	Links           map[string]string  `json:"Links"`
}

type ActivityElement struct {
	ID                 string                `json:"Id"`
	Name               string                `json:"Name"`
	Started            *time.Time            `json:"Started"`
	Ended              *time.Time            `json:"Ended"`
	Status             string                `json:"Status"`
	Children           []*ActivityElement    `json:"Children"`
	ShowAtSummaryLevel bool                  `json:"ShowAtSummaryLevel"`
	LogElements        []*ActivityLogElement `json:"LogElements"`
}

type ActivityLogElement struct {
	OccurredAt  time.Time `json:"OccurredAt"`
	Category    string    `json:"Category"`
	MessageText string    `json:"MessageText"`
	Number      int       `json:"Number"`
}

type TaskProgress struct {
	ProgressPercentage     int    `json:"ProgressPercentage"`
	EstimatedTimeRemaining string `json:"EstimatedTimeRemaining"`
}
