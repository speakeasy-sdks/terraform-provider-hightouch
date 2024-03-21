// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/internal/utils"
)

// TriggerRunIDGraphInput - The input of a trigger action to run IDR.
type TriggerRunIDGraphInput struct {
	// Whether to resync the entire Identity Graph or process incrementally.
	FullRerun *bool `default:"false" json:"fullRerun"`
}

func (t TriggerRunIDGraphInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TriggerRunIDGraphInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TriggerRunIDGraphInput) GetFullRerun() *bool {
	if o == nil {
		return nil
	}
	return o.FullRerun
}