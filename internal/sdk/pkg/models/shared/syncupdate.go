// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"hightouch/internal/sdk/pkg/utils"
)

type SyncUpdateScheduleScheduleType string

const (
	SyncUpdateScheduleScheduleTypeIntervalSchedule   SyncUpdateScheduleScheduleType = "IntervalSchedule"
	SyncUpdateScheduleScheduleTypeCronSchedule       SyncUpdateScheduleScheduleType = "CronSchedule"
	SyncUpdateScheduleScheduleTypeVisualCronSchedule SyncUpdateScheduleScheduleType = "VisualCronSchedule"
	SyncUpdateScheduleScheduleTypeDBTSchedule        SyncUpdateScheduleScheduleType = "DBTSchedule"
)

type SyncUpdateScheduleSchedule struct {
	IntervalSchedule   *IntervalSchedule
	CronSchedule       *CronSchedule
	VisualCronSchedule *VisualCronSchedule
	DBTSchedule        *DBTSchedule

	Type SyncUpdateScheduleScheduleType
}

func CreateSyncUpdateScheduleScheduleIntervalSchedule(intervalSchedule IntervalSchedule) SyncUpdateScheduleSchedule {
	typ := SyncUpdateScheduleScheduleTypeIntervalSchedule

	return SyncUpdateScheduleSchedule{
		IntervalSchedule: &intervalSchedule,
		Type:             typ,
	}
}

func CreateSyncUpdateScheduleScheduleCronSchedule(cronSchedule CronSchedule) SyncUpdateScheduleSchedule {
	typ := SyncUpdateScheduleScheduleTypeCronSchedule

	return SyncUpdateScheduleSchedule{
		CronSchedule: &cronSchedule,
		Type:         typ,
	}
}

func CreateSyncUpdateScheduleScheduleVisualCronSchedule(visualCronSchedule VisualCronSchedule) SyncUpdateScheduleSchedule {
	typ := SyncUpdateScheduleScheduleTypeVisualCronSchedule

	return SyncUpdateScheduleSchedule{
		VisualCronSchedule: &visualCronSchedule,
		Type:               typ,
	}
}

func CreateSyncUpdateScheduleScheduleDBTSchedule(dbtSchedule DBTSchedule) SyncUpdateScheduleSchedule {
	typ := SyncUpdateScheduleScheduleTypeDBTSchedule

	return SyncUpdateScheduleSchedule{
		DBTSchedule: &dbtSchedule,
		Type:        typ,
	}
}

func (u *SyncUpdateScheduleSchedule) UnmarshalJSON(data []byte) error {

	intervalSchedule := new(IntervalSchedule)
	if err := utils.UnmarshalJSON(data, &intervalSchedule, "", true, true); err == nil {
		u.IntervalSchedule = intervalSchedule
		u.Type = SyncUpdateScheduleScheduleTypeIntervalSchedule
		return nil
	}

	cronSchedule := new(CronSchedule)
	if err := utils.UnmarshalJSON(data, &cronSchedule, "", true, true); err == nil {
		u.CronSchedule = cronSchedule
		u.Type = SyncUpdateScheduleScheduleTypeCronSchedule
		return nil
	}

	visualCronSchedule := new(VisualCronSchedule)
	if err := utils.UnmarshalJSON(data, &visualCronSchedule, "", true, true); err == nil {
		u.VisualCronSchedule = visualCronSchedule
		u.Type = SyncUpdateScheduleScheduleTypeVisualCronSchedule
		return nil
	}

	dbtSchedule := new(DBTSchedule)
	if err := utils.UnmarshalJSON(data, &dbtSchedule, "", true, true); err == nil {
		u.DBTSchedule = dbtSchedule
		u.Type = SyncUpdateScheduleScheduleTypeDBTSchedule
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SyncUpdateScheduleSchedule) MarshalJSON() ([]byte, error) {
	if u.IntervalSchedule != nil {
		return utils.MarshalJSON(u.IntervalSchedule, "", true)
	}

	if u.CronSchedule != nil {
		return utils.MarshalJSON(u.CronSchedule, "", true)
	}

	if u.VisualCronSchedule != nil {
		return utils.MarshalJSON(u.VisualCronSchedule, "", true)
	}

	if u.DBTSchedule != nil {
		return utils.MarshalJSON(u.DBTSchedule, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// SyncUpdateSchedule - The scheduling configuration. It can be triggerd based on several ways:
//
// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
//
// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
//
// Visual: the sync will be trigged based a visual cron configuration on UI
//
// DBT-cloud: the sync will be trigged based on a dbt cloud job
type SyncUpdateSchedule struct {
	Schedule SyncUpdateScheduleSchedule `json:"schedule"`
	Type     string                     `json:"type"`
}

func (o *SyncUpdateSchedule) GetSchedule() SyncUpdateScheduleSchedule {
	if o == nil {
		return SyncUpdateScheduleSchedule{}
	}
	return o.Schedule
}

func (o *SyncUpdateSchedule) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

// SyncUpdate - The input for updating a Sync
type SyncUpdate struct {
	// The sync's configuration. This specifies how data is mapped, among other
	// configuration.
	//
	// The schema depends on the destination type.
	//
	// Consumers should NOT make assumptions on the contents of the
	// configuration. It may change as Hightouch updates its internal code.
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	// Whether the sync has been disabled by the user.
	Disabled *bool `json:"disabled,omitempty"`
	// The scheduling configuration. It can be triggerd based on several ways:
	//
	// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
	//
	// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
	//
	// Visual: the sync will be trigged based a visual cron configuration on UI
	//
	// DBT-cloud: the sync will be trigged based on a dbt cloud job
	Schedule *SyncUpdateSchedule `json:"schedule,omitempty"`
}

func (o *SyncUpdate) GetConfiguration() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Configuration
}

func (o *SyncUpdate) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *SyncUpdate) GetSchedule() *SyncUpdateSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}
