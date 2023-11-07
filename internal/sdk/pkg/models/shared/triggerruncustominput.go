// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"hightouch/v2/internal/sdk/pkg/utils"
)

// TriggerRunCustomInput - The input of a trigger action to run syncs based on sync ID, slug or other filters
type TriggerRunCustomInput struct {
	// Whether to resync all the rows in the query (i.e. ignoring previously
	// synced rows).
	FullResync *bool `default:"false" json:"fullResync"`
	// Trigger run based on sync id
	SyncID *string `json:"syncId,omitempty"`
	// Trigger run based on sync slug
	SyncSlug *string `json:"syncSlug,omitempty"`
}

func (t TriggerRunCustomInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TriggerRunCustomInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TriggerRunCustomInput) GetFullResync() *bool {
	if o == nil {
		return nil
	}
	return o.FullResync
}

func (o *TriggerRunCustomInput) GetSyncID() *string {
	if o == nil {
		return nil
	}
	return o.SyncID
}

func (o *TriggerRunCustomInput) GetSyncSlug() *string {
	if o == nil {
		return nil
	}
	return o.SyncSlug
}
