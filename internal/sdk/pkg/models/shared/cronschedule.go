// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CronSchedule struct {
	Expression string `json:"expression"`
}

func (o *CronSchedule) GetExpression() string {
	if o == nil {
		return ""
	}
	return o.Expression
}
