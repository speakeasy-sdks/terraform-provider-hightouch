// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/de-tf-providers/terraform-provider-hightouch/v2/internal/sdk/pkg/models/shared"
	"github.com/de-tf-providers/terraform-provider-hightouch/v2/internal/sdk/pkg/utils"
	"net/http"
)

type TriggerRunCustomResponseBodyType string

const (
	TriggerRunCustomResponseBodyTypeTriggerRunOutput  TriggerRunCustomResponseBodyType = "TriggerRunOutput"
	TriggerRunCustomResponseBodyTypeValidateErrorJSON TriggerRunCustomResponseBodyType = "ValidateErrorJSON"
)

type TriggerRunCustomResponseBody struct {
	TriggerRunOutput  *shared.TriggerRunOutput
	ValidateErrorJSON *shared.ValidateErrorJSON

	Type TriggerRunCustomResponseBodyType
}

func CreateTriggerRunCustomResponseBodyTriggerRunOutput(triggerRunOutput shared.TriggerRunOutput) TriggerRunCustomResponseBody {
	typ := TriggerRunCustomResponseBodyTypeTriggerRunOutput

	return TriggerRunCustomResponseBody{
		TriggerRunOutput: &triggerRunOutput,
		Type:             typ,
	}
}

func CreateTriggerRunCustomResponseBodyValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) TriggerRunCustomResponseBody {
	typ := TriggerRunCustomResponseBodyTypeValidateErrorJSON

	return TriggerRunCustomResponseBody{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func (u *TriggerRunCustomResponseBody) UnmarshalJSON(data []byte) error {

	triggerRunOutput := new(shared.TriggerRunOutput)
	if err := utils.UnmarshalJSON(data, &triggerRunOutput, "", true, true); err == nil {
		u.TriggerRunOutput = triggerRunOutput
		u.Type = TriggerRunCustomResponseBodyTypeTriggerRunOutput
		return nil
	}

	validateErrorJSON := new(shared.ValidateErrorJSON)
	if err := utils.UnmarshalJSON(data, &validateErrorJSON, "", true, true); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = TriggerRunCustomResponseBodyTypeValidateErrorJSON
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TriggerRunCustomResponseBody) MarshalJSON() ([]byte, error) {
	if u.TriggerRunOutput != nil {
		return utils.MarshalJSON(u.TriggerRunOutput, "", true)
	}

	if u.ValidateErrorJSON != nil {
		return utils.MarshalJSON(u.ValidateErrorJSON, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type TriggerRunCustomResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
	// Ok
	OneOf *TriggerRunCustomResponseBody
}

func (o *TriggerRunCustomResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TriggerRunCustomResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TriggerRunCustomResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TriggerRunCustomResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}

func (o *TriggerRunCustomResponse) GetOneOf() *TriggerRunCustomResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
