// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type IntervalSchedule struct {
	Interval Interval `json:"interval"`
}

func (o *IntervalSchedule) GetInterval() Interval {
	if o == nil {
		return Interval{}
	}
	return o.Interval
}