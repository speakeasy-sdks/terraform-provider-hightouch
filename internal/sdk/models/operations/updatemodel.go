// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/internal/utils"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/models/shared"
	"net/http"
)

type UpdateModelRequest struct {
	ModelUpdate shared.ModelUpdate `request:"mediaType=application/json"`
	// The model's ID
	ModelID float64 `pathParam:"style=simple,explode=false,name=modelId"`
}

func (o *UpdateModelRequest) GetModelUpdate() shared.ModelUpdate {
	if o == nil {
		return shared.ModelUpdate{}
	}
	return o.ModelUpdate
}

func (o *UpdateModelRequest) GetModelID() float64 {
	if o == nil {
		return 0.0
	}
	return o.ModelID
}

type UpdateModelResponseBodyType string

const (
	UpdateModelResponseBodyTypeModel               UpdateModelResponseBodyType = "Model"
	UpdateModelResponseBodyTypeValidateErrorJSON   UpdateModelResponseBodyType = "ValidateErrorJSON"
	UpdateModelResponseBodyTypeInternalServerError UpdateModelResponseBodyType = "InternalServerError"
)

// UpdateModelResponseBody - Ok
type UpdateModelResponseBody struct {
	Model               *shared.Model
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type UpdateModelResponseBodyType
}

func CreateUpdateModelResponseBodyModel(model shared.Model) UpdateModelResponseBody {
	typ := UpdateModelResponseBodyTypeModel

	return UpdateModelResponseBody{
		Model: &model,
		Type:  typ,
	}
}

func CreateUpdateModelResponseBodyValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) UpdateModelResponseBody {
	typ := UpdateModelResponseBodyTypeValidateErrorJSON

	return UpdateModelResponseBody{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateUpdateModelResponseBodyInternalServerError(internalServerError shared.InternalServerError) UpdateModelResponseBody {
	typ := UpdateModelResponseBodyTypeInternalServerError

	return UpdateModelResponseBody{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *UpdateModelResponseBody) UnmarshalJSON(data []byte) error {

	validateErrorJSON := shared.ValidateErrorJSON{}
	if err := utils.UnmarshalJSON(data, &validateErrorJSON, "", true, true); err == nil {
		u.ValidateErrorJSON = &validateErrorJSON
		u.Type = UpdateModelResponseBodyTypeValidateErrorJSON
		return nil
	}

	model := shared.Model{}
	if err := utils.UnmarshalJSON(data, &model, "", true, true); err == nil {
		u.Model = &model
		u.Type = UpdateModelResponseBodyTypeModel
		return nil
	}

	internalServerError := shared.InternalServerError("")
	if err := utils.UnmarshalJSON(data, &internalServerError, "", true, true); err == nil {
		u.InternalServerError = &internalServerError
		u.Type = UpdateModelResponseBodyTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdateModelResponseBody) MarshalJSON() ([]byte, error) {
	if u.Model != nil {
		return utils.MarshalJSON(u.Model, "", true)
	}

	if u.ValidateErrorJSON != nil {
		return utils.MarshalJSON(u.ValidateErrorJSON, "", true)
	}

	if u.InternalServerError != nil {
		return utils.MarshalJSON(u.InternalServerError, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UpdateModelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Something went wrong
	InternalServerError *shared.InternalServerError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
	// Ok
	OneOf *UpdateModelResponseBody
}

func (o *UpdateModelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateModelResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateModelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateModelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateModelResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}

func (o *UpdateModelResponse) GetOneOf() *UpdateModelResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
