// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

type VisualCronScheduleExpressions struct {
	// Construct a type with a set of properties K of type T
	Days RecordDayBooleanOrUndefined `json:"days"`
	Time string                      `json:"time"`
}

type VisualCronSchedule struct {
	Expressions []VisualCronScheduleExpressions `json:"expressions"`

	AdditionalProperties interface{} `json:"-"`
}
type _VisualCronSchedule VisualCronSchedule

func (c *VisualCronSchedule) UnmarshalJSON(bs []byte) error {
	data := _VisualCronSchedule{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = VisualCronSchedule(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "expressions")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c VisualCronSchedule) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_VisualCronSchedule(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}
