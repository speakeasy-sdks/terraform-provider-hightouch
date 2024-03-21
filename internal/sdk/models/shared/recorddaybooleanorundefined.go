// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RecordDayBooleanOrUndefined - Construct a type with a set of properties K of type T
type RecordDayBooleanOrUndefined struct {
	Friday    *bool `json:"friday,omitempty"`
	Monday    *bool `json:"monday,omitempty"`
	Saturday  *bool `json:"saturday,omitempty"`
	Sunday    *bool `json:"sunday,omitempty"`
	Thursday  *bool `json:"thursday,omitempty"`
	Tuesday   *bool `json:"tuesday,omitempty"`
	Wednesday *bool `json:"wednesday,omitempty"`
}

func (o *RecordDayBooleanOrUndefined) GetFriday() *bool {
	if o == nil {
		return nil
	}
	return o.Friday
}

func (o *RecordDayBooleanOrUndefined) GetMonday() *bool {
	if o == nil {
		return nil
	}
	return o.Monday
}

func (o *RecordDayBooleanOrUndefined) GetSaturday() *bool {
	if o == nil {
		return nil
	}
	return o.Saturday
}

func (o *RecordDayBooleanOrUndefined) GetSunday() *bool {
	if o == nil {
		return nil
	}
	return o.Sunday
}

func (o *RecordDayBooleanOrUndefined) GetThursday() *bool {
	if o == nil {
		return nil
	}
	return o.Thursday
}

func (o *RecordDayBooleanOrUndefined) GetTuesday() *bool {
	if o == nil {
		return nil
	}
	return o.Tuesday
}

func (o *RecordDayBooleanOrUndefined) GetWednesday() *bool {
	if o == nil {
		return nil
	}
	return o.Wednesday
}