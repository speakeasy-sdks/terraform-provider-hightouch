// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/pkg/models/shared"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/pkg/utils"
	"net/http"
)

type UpdateSourceRequest struct {
	SourceUpdate shared.SourceUpdate `request:"mediaType=application/json"`
	// The source's ID
	SourceID float64 `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *UpdateSourceRequest) GetSourceUpdate() shared.SourceUpdate {
	if o == nil {
		return shared.SourceUpdate{}
	}
	return o.SourceUpdate
}

func (o *UpdateSourceRequest) GetSourceID() float64 {
	if o == nil {
		return 0.0
	}
	return o.SourceID
}

type UpdateSourceResponseBodyType string

const (
	UpdateSourceResponseBodyTypeSource              UpdateSourceResponseBodyType = "Source"
	UpdateSourceResponseBodyTypeValidateErrorJSON   UpdateSourceResponseBodyType = "ValidateErrorJSON"
	UpdateSourceResponseBodyTypeInternalServerError UpdateSourceResponseBodyType = "InternalServerError"
)

// UpdateSourceResponseBody - Ok
type UpdateSourceResponseBody struct {
	Source              *shared.Source
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type UpdateSourceResponseBodyType
}

func CreateUpdateSourceResponseBodySource(source shared.Source) UpdateSourceResponseBody {
	typ := UpdateSourceResponseBodyTypeSource

	return UpdateSourceResponseBody{
		Source: &source,
		Type:   typ,
	}
}

func CreateUpdateSourceResponseBodyValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) UpdateSourceResponseBody {
	typ := UpdateSourceResponseBodyTypeValidateErrorJSON

	return UpdateSourceResponseBody{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateUpdateSourceResponseBodyInternalServerError(internalServerError shared.InternalServerError) UpdateSourceResponseBody {
	typ := UpdateSourceResponseBodyTypeInternalServerError

	return UpdateSourceResponseBody{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *UpdateSourceResponseBody) UnmarshalJSON(data []byte) error {

	validateErrorJSON := new(shared.ValidateErrorJSON)
	if err := utils.UnmarshalJSON(data, &validateErrorJSON, "", true, true); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = UpdateSourceResponseBodyTypeValidateErrorJSON
		return nil
	}

	source := new(shared.Source)
	if err := utils.UnmarshalJSON(data, &source, "", true, true); err == nil {
		u.Source = source
		u.Type = UpdateSourceResponseBodyTypeSource
		return nil
	}

	internalServerError := new(shared.InternalServerError)
	if err := utils.UnmarshalJSON(data, &internalServerError, "", true, true); err == nil {
		u.InternalServerError = internalServerError
		u.Type = UpdateSourceResponseBodyTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdateSourceResponseBody) MarshalJSON() ([]byte, error) {
	if u.Source != nil {
		return utils.MarshalJSON(u.Source, "", true)
	}

	if u.ValidateErrorJSON != nil {
		return utils.MarshalJSON(u.ValidateErrorJSON, "", true)
	}

	if u.InternalServerError != nil {
		return utils.MarshalJSON(u.InternalServerError, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UpdateSourceResponse struct {
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
	OneOf *UpdateSourceResponseBody
}

func (o *UpdateSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSourceResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSourceResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}

func (o *UpdateSourceResponse) GetOneOf() *UpdateSourceResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
