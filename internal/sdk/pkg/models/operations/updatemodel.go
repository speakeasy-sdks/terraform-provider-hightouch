// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateModelRequest struct {
	ModelUpdate shared.ModelUpdate `request:"mediaType=application/json"`
	// The model's ID
	ModelID float64 `pathParam:"style=simple,explode=false,name=modelId"`
}

type UpdateModel200ApplicationJSONType string

const (
	UpdateModel200ApplicationJSONTypeModel               UpdateModel200ApplicationJSONType = "Model"
	UpdateModel200ApplicationJSONTypeValidateErrorJSON   UpdateModel200ApplicationJSONType = "ValidateErrorJSON"
	UpdateModel200ApplicationJSONTypeInternalServerError UpdateModel200ApplicationJSONType = "InternalServerError"
)

type UpdateModel200ApplicationJSON struct {
	Model               *shared.Model
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type UpdateModel200ApplicationJSONType
}

func CreateUpdateModel200ApplicationJSONModel(model shared.Model) UpdateModel200ApplicationJSON {
	typ := UpdateModel200ApplicationJSONTypeModel

	return UpdateModel200ApplicationJSON{
		Model: &model,
		Type:  typ,
	}
}

func CreateUpdateModel200ApplicationJSONValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) UpdateModel200ApplicationJSON {
	typ := UpdateModel200ApplicationJSONTypeValidateErrorJSON

	return UpdateModel200ApplicationJSON{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateUpdateModel200ApplicationJSONInternalServerError(internalServerError shared.InternalServerError) UpdateModel200ApplicationJSON {
	typ := UpdateModel200ApplicationJSONTypeInternalServerError

	return UpdateModel200ApplicationJSON{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *UpdateModel200ApplicationJSON) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	internalServerError := new(shared.InternalServerError)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&internalServerError); err == nil {
		u.InternalServerError = internalServerError
		u.Type = UpdateModel200ApplicationJSONTypeInternalServerError
		return nil
	}

	validateErrorJSON := new(shared.ValidateErrorJSON)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&validateErrorJSON); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = UpdateModel200ApplicationJSONTypeValidateErrorJSON
		return nil
	}

	model := new(shared.Model)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&model); err == nil {
		u.Model = model
		u.Type = UpdateModel200ApplicationJSONTypeModel
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdateModel200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.InternalServerError != nil {
		return json.Marshal(u.InternalServerError)
	}

	if u.ValidateErrorJSON != nil {
		return json.Marshal(u.ValidateErrorJSON)
	}

	if u.Model != nil {
		return json.Marshal(u.Model)
	}

	return nil, nil
}

type UpdateModelResponse struct {
	ContentType string
	// Something went wrong
	InternalServerError *shared.InternalServerError
	StatusCode          int
	RawResponse         *http.Response
	// Ok
	UpdateModel200ApplicationJSONOneOf *UpdateModel200ApplicationJSON
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
