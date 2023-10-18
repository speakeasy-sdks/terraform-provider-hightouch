// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"hightouch/internal/sdk/pkg/utils"
)

// TriggerRunInput - The input of a trigger action to run syncs
type TriggerRunInput struct {
	// Whether to resync all the rows in the query (i.e. ignoring previously
	// synced rows).
	FullResync *bool `default:"false" json:"fullResync"`
}

func (t TriggerRunInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TriggerRunInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TriggerRunInput) GetFullResync() *bool {
	if o == nil {
		return nil
	}
	return o.FullResync
}
