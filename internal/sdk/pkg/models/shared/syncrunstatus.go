// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SyncRunStatus - The status of sync runs
type SyncRunStatus string

const (
	SyncRunStatusCancelled   SyncRunStatus = "cancelled"
	SyncRunStatusFailed      SyncRunStatus = "failed"
	SyncRunStatusQueued      SyncRunStatus = "queued"
	SyncRunStatusSuccess     SyncRunStatus = "success"
	SyncRunStatusWarning     SyncRunStatus = "warning"
	SyncRunStatusQuerying    SyncRunStatus = "querying"
	SyncRunStatusProcessing  SyncRunStatus = "processing"
	SyncRunStatusReporting   SyncRunStatus = "reporting"
	SyncRunStatusInterrupted SyncRunStatus = "interrupted"
)

func (e SyncRunStatus) ToPointer() *SyncRunStatus {
	return &e
}

func (e *SyncRunStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cancelled":
		fallthrough
	case "failed":
		fallthrough
	case "queued":
		fallthrough
	case "success":
		fallthrough
	case "warning":
		fallthrough
	case "querying":
		fallthrough
	case "processing":
		fallthrough
	case "reporting":
		fallthrough
	case "interrupted":
		*e = SyncRunStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SyncRunStatus: %v", v)
	}
}
