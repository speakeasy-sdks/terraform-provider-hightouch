// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/internal/utils"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/models/shared"
	"net/http"
)

type UpdateDestinationRequest struct {
	DestinationUpdate shared.DestinationUpdate `request:"mediaType=application/json"`
	// The destination's ID
	DestinationID float64 `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *UpdateDestinationRequest) GetDestinationUpdate() shared.DestinationUpdate {
	if o == nil {
		return shared.DestinationUpdate{}
	}
	return o.DestinationUpdate
}

func (o *UpdateDestinationRequest) GetDestinationID() float64 {
	if o == nil {
		return 0.0
	}
	return o.DestinationID
}

type UpdateDestinationResponseBodyType string

const (
	UpdateDestinationResponseBodyTypeDestination         UpdateDestinationResponseBodyType = "Destination"
	UpdateDestinationResponseBodyTypeValidateErrorJSON   UpdateDestinationResponseBodyType = "ValidateErrorJSON"
	UpdateDestinationResponseBodyTypeInternalServerError UpdateDestinationResponseBodyType = "InternalServerError"
)

// UpdateDestinationResponseBody - Ok
type UpdateDestinationResponseBody struct {
	Destination         *shared.Destination
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type UpdateDestinationResponseBodyType
}

func CreateUpdateDestinationResponseBodyDestination(destination shared.Destination) UpdateDestinationResponseBody {
	typ := UpdateDestinationResponseBodyTypeDestination

	return UpdateDestinationResponseBody{
		Destination: &destination,
		Type:        typ,
	}
}

func CreateUpdateDestinationResponseBodyValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) UpdateDestinationResponseBody {
	typ := UpdateDestinationResponseBodyTypeValidateErrorJSON

	return UpdateDestinationResponseBody{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateUpdateDestinationResponseBodyInternalServerError(internalServerError shared.InternalServerError) UpdateDestinationResponseBody {
	typ := UpdateDestinationResponseBodyTypeInternalServerError

	return UpdateDestinationResponseBody{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *UpdateDestinationResponseBody) UnmarshalJSON(data []byte) error {

	validateErrorJSON := shared.ValidateErrorJSON{}
	if err := utils.UnmarshalJSON(data, &validateErrorJSON, "", true, true); err == nil {
		u.ValidateErrorJSON = &validateErrorJSON
		u.Type = UpdateDestinationResponseBodyTypeValidateErrorJSON
		return nil
	}

	destination := shared.Destination{}
	if err := utils.UnmarshalJSON(data, &destination, "", true, true); err == nil {
		u.Destination = &destination
		u.Type = UpdateDestinationResponseBodyTypeDestination
		return nil
	}

	internalServerError := shared.InternalServerError("")
	if err := utils.UnmarshalJSON(data, &internalServerError, "", true, true); err == nil {
		u.InternalServerError = &internalServerError
		u.Type = UpdateDestinationResponseBodyTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdateDestinationResponseBody) MarshalJSON() ([]byte, error) {
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

type UpdateDestinationResponse struct {
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
	OneOf *UpdateDestinationResponseBody
}

func (o *UpdateDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDestinationResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateDestinationResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}

func (o *UpdateDestinationResponse) GetOneOf() *UpdateDestinationResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
