// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/de-tf-providers/terraform-provider-hightouch/v2/internal/sdk/pkg/models/shared"
	"github.com/de-tf-providers/terraform-provider-hightouch/v2/internal/sdk/pkg/utils"
	"net/http"
)

type CreateDestinationResponseBodyType string

const (
	CreateDestinationResponseBodyTypeDestination         CreateDestinationResponseBodyType = "Destination"
	CreateDestinationResponseBodyTypeValidateErrorJSON   CreateDestinationResponseBodyType = "ValidateErrorJSON"
	CreateDestinationResponseBodyTypeInternalServerError CreateDestinationResponseBodyType = "InternalServerError"
)

type CreateDestinationResponseBody struct {
	Destination         *shared.Destination
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type CreateDestinationResponseBodyType
}

func CreateCreateDestinationResponseBodyDestination(destination shared.Destination) CreateDestinationResponseBody {
	typ := CreateDestinationResponseBodyTypeDestination

	return CreateDestinationResponseBody{
		Destination: &destination,
		Type:        typ,
	}
}

func CreateCreateDestinationResponseBodyValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) CreateDestinationResponseBody {
	typ := CreateDestinationResponseBodyTypeValidateErrorJSON

	return CreateDestinationResponseBody{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateCreateDestinationResponseBodyInternalServerError(internalServerError shared.InternalServerError) CreateDestinationResponseBody {
	typ := CreateDestinationResponseBodyTypeInternalServerError

	return CreateDestinationResponseBody{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *CreateDestinationResponseBody) UnmarshalJSON(data []byte) error {

	validateErrorJSON := new(shared.ValidateErrorJSON)
	if err := utils.UnmarshalJSON(data, &validateErrorJSON, "", true, true); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = CreateDestinationResponseBodyTypeValidateErrorJSON
		return nil
	}

	destination := new(shared.Destination)
	if err := utils.UnmarshalJSON(data, &destination, "", true, true); err == nil {
		u.Destination = destination
		u.Type = CreateDestinationResponseBodyTypeDestination
		return nil
	}

	internalServerError := new(shared.InternalServerError)
	if err := utils.UnmarshalJSON(data, &internalServerError, "", true, true); err == nil {
		u.InternalServerError = internalServerError
		u.Type = CreateDestinationResponseBodyTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateDestinationResponseBody) MarshalJSON() ([]byte, error) {
	if u.Destination != nil {
		return utils.MarshalJSON(u.Destination, "", true)
	}

	if u.ValidateErrorJSON != nil {
		return utils.MarshalJSON(u.ValidateErrorJSON, "", true)
	}

	if u.InternalServerError != nil {
		return utils.MarshalJSON(u.InternalServerError, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateDestinationResponse struct {
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
	OneOf *CreateDestinationResponseBody
}

func (o *CreateDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDestinationResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *CreateDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateDestinationResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}

func (o *CreateDestinationResponse) GetOneOf() *CreateDestinationResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
