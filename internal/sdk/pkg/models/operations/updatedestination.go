// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateDestinationRequest struct {
	DestinationUpdate shared.DestinationUpdate `request:"mediaType=application/json"`
	// The destination's ID
	DestinationID float64 `pathParam:"style=simple,explode=false,name=destinationId"`
}

type UpdateDestination200ApplicationJSONType string

const (
	UpdateDestination200ApplicationJSONTypeDestination         UpdateDestination200ApplicationJSONType = "Destination"
	UpdateDestination200ApplicationJSONTypeValidateErrorJSON   UpdateDestination200ApplicationJSONType = "ValidateErrorJSON"
	UpdateDestination200ApplicationJSONTypeInternalServerError UpdateDestination200ApplicationJSONType = "InternalServerError"
)

type UpdateDestination200ApplicationJSON struct {
	Destination         *shared.Destination
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type UpdateDestination200ApplicationJSONType
}

func CreateUpdateDestination200ApplicationJSONDestination(destination shared.Destination) UpdateDestination200ApplicationJSON {
	typ := UpdateDestination200ApplicationJSONTypeDestination

	return UpdateDestination200ApplicationJSON{
		Destination: &destination,
		Type:        typ,
	}
}

func CreateUpdateDestination200ApplicationJSONValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) UpdateDestination200ApplicationJSON {
	typ := UpdateDestination200ApplicationJSONTypeValidateErrorJSON

	return UpdateDestination200ApplicationJSON{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateUpdateDestination200ApplicationJSONInternalServerError(internalServerError shared.InternalServerError) UpdateDestination200ApplicationJSON {
	typ := UpdateDestination200ApplicationJSONTypeInternalServerError

	return UpdateDestination200ApplicationJSON{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *UpdateDestination200ApplicationJSON) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	validateErrorJSON := new(shared.ValidateErrorJSON)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&validateErrorJSON); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = UpdateDestination200ApplicationJSONTypeValidateErrorJSON
		return nil
	}

	internalServerError := new(shared.InternalServerError)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&internalServerError); err == nil {
		u.InternalServerError = internalServerError
		u.Type = UpdateDestination200ApplicationJSONTypeInternalServerError
		return nil
	}

	destination := new(shared.Destination)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destination); err == nil {
		u.Destination = destination
		u.Type = UpdateDestination200ApplicationJSONTypeDestination
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdateDestination200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.ValidateErrorJSON != nil {
		return json.Marshal(u.ValidateErrorJSON)
	}

	if u.InternalServerError != nil {
		return json.Marshal(u.InternalServerError)
	}

	if u.Destination != nil {
		return json.Marshal(u.Destination)
	}

	return nil, nil
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
	// Ok
	UpdateDestination200ApplicationJSONOneOf *UpdateDestination200ApplicationJSON
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
