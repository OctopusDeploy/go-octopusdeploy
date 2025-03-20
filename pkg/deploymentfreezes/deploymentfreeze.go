package deploymentfreezes

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type TenantProjectEnvironment struct {
	TenantId      string `json:"TenantId"`
	ProjectId     string `json:"ProjectId"`
	EnvironmentId string `json:"EnvironmentId"`
}

type RecurringScheduleType string

const (
	Daily    RecurringScheduleType = "Daily"
	Weekly   RecurringScheduleType = "Weekly"
	Monthly  RecurringScheduleType = "Monthly"
	Annually RecurringScheduleType = "Annually"
)

type RecurringScheduleEndType string

const (
	Never            RecurringScheduleEndType = "Never"
	OnDate           RecurringScheduleEndType = "OnDate"
	AfterOccurrences RecurringScheduleEndType = "AfterOccurrences"
)

type MonthlyScheduleType string

const (
	DayOfMonth  MonthlyScheduleType = "DayOfMonth"
	DateOfMonth MonthlyScheduleType = "DateOfMonth"
)

type RecurringSchedule struct {
	Type                RecurringScheduleType    `json:"Type"`
	Unit                int                      `json:"Unit"`
	EndType             RecurringScheduleEndType `json:"EndType"`
	EndOnDate           *time.Time               `json:"EndOnDate,omitempty"`
	EndAfterOccurrences int                      `json:"EndAfterOccurrences,omitempty"`
	MonthlyScheduleType string                   `json:"MonthlyScheduleType,omitempty"`
	DateOfMonth         string                   `json:"DateOfMonth,omitempty"`
	DayNumberOfMonth    string                   `json:"DayNumberOfMonth,omitempty"`
	DaysOfWeek          []string                 `json:"DaysOfWeek,omitempty"`
	DayOfWeek           string                   `json:"DayOfWeek,omitempty"`
}

type DeploymentFreezes struct {
	DeploymentFreezes []DeploymentFreeze `json:"DeploymentFreezes"`
	Count             int                `json:"Count"`
}

type DeploymentFreeze struct {
	Name                          string                     `json:"Name" validate:"required"`
	Start                         *time.Time                 `json:"Start" validate:"required"`
	End                           *time.Time                 `json:"End" validate:"required"`
	ProjectEnvironmentScope       map[string][]string        `json:"ProjectEnvironmentScope,omitempty"`
	TenantProjectEnvironmentScope []TenantProjectEnvironment `json:"TenantProjectEnvironmentScope,omitempty"`
	RecurringSchedule             *RecurringSchedule         `json:"RecurringSchedule,omitempty"`
	OwnerId                       string                     `json:"OwnerId,omitempty"`

	resources.Resource
}

func (d *DeploymentFreeze) GetName() string {
	return d.Name
}

func (d *DeploymentFreeze) SetName(name string) {
	d.Name = name
}
