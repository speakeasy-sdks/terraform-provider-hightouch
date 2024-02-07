// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/pkg/models/shared"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/pkg/utils"
	"net/http"
)

type CreateSourceResponseBodyType string

const (
	CreateSourceResponseBodyTypeSource              CreateSourceResponseBodyType = "Source"
	CreateSourceResponseBodyTypeValidateErrorJSON   CreateSourceResponseBodyType = "ValidateErrorJSON"
	CreateSourceResponseBodyTypeInternalServerError CreateSourceResponseBodyType = "InternalServerError"
)

// CreateSourceResponseBody - Ok
type CreateSourceResponseBody struct {
	Source              *shared.Source
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type CreateSourceResponseBodyType
}

func CreateCreateSourceResponseBodySource(source shared.Source) CreateSourceResponseBody {
	typ := CreateSourceResponseBodyTypeSource

	return CreateSourceResponseBody{
		Source: &source,
		Type:   typ,
	}
}

func CreateCreateSourceResponseBodyValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) CreateSourceResponseBody {
	typ := CreateSourceResponseBodyTypeValidateErrorJSON

	return CreateSourceResponseBody{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateCreateSourceResponseBodyInternalServerError(internalServerError shared.InternalServerError) CreateSourceResponseBody {
	typ := CreateSourceResponseBodyTypeInternalServerError

	return CreateSourceResponseBody{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *CreateSourceResponseBody) UnmarshalJSON(data []byte) error {

	validateErrorJSON := new(shared.ValidateErrorJSON)
	if err := utils.UnmarshalJSON(data, &validateErrorJSON, "", true, true); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = CreateSourceResponseBodyTypeValidateErrorJSON
		return nil
	}

	source := new(shared.Source)
	if err := utils.UnmarshalJSON(data, &source, "", true, true); err == nil {
		u.Source = source
		u.Type = CreateSourceResponseBodyTypeSource
		return nil
	}

	internalServerError := new(shared.InternalServerError)
	if err := utils.UnmarshalJSON(data, &internalServerError, "", true, true); err == nil {
		u.InternalServerError = internalServerError
		u.Type = CreateSourceResponseBodyTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateSourceResponseBody) MarshalJSON() ([]byte, error) {
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

type CreateSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Something went wrong
	InternalServerError *shared.InternalServerError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Conflict
	ValidateErrorJSON *shared.ValidateErrorJSON
	// Ok
	OneOf *CreateSourceResponseBody
}

func (o *CreateSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSourceResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *CreateSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSourceResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}

func (o *CreateSourceResponse) GetOneOf() *CreateSourceResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
