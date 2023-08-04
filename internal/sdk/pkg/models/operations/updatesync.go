// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSyncSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type UpdateSyncRequest struct {
	SyncUpdate shared.SyncUpdate `request:"mediaType=application/json"`
	// The sync's ID
	SyncID float64 `pathParam:"style=simple,explode=false,name=syncId"`
}

type UpdateSync200ApplicationJSONType string

const (
	UpdateSync200ApplicationJSONTypeSync                UpdateSync200ApplicationJSONType = "Sync"
	UpdateSync200ApplicationJSONTypeValidateErrorJSON   UpdateSync200ApplicationJSONType = "ValidateErrorJSON"
	UpdateSync200ApplicationJSONTypeInternalServerError UpdateSync200ApplicationJSONType = "InternalServerError"
)

type UpdateSync200ApplicationJSON struct {
	Sync                *shared.Sync
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type UpdateSync200ApplicationJSONType
}

func CreateUpdateSync200ApplicationJSONSync(sync shared.Sync) UpdateSync200ApplicationJSON {
	typ := UpdateSync200ApplicationJSONTypeSync

	return UpdateSync200ApplicationJSON{
		Sync: &sync,
		Type: typ,
	}
}

func CreateUpdateSync200ApplicationJSONValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) UpdateSync200ApplicationJSON {
	typ := UpdateSync200ApplicationJSONTypeValidateErrorJSON

	return UpdateSync200ApplicationJSON{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateUpdateSync200ApplicationJSONInternalServerError(internalServerError shared.InternalServerError) UpdateSync200ApplicationJSON {
	typ := UpdateSync200ApplicationJSONTypeInternalServerError

	return UpdateSync200ApplicationJSON{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *UpdateSync200ApplicationJSON) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sync := new(shared.Sync)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sync); err == nil {
		u.Sync = sync
		u.Type = UpdateSync200ApplicationJSONTypeSync
		return nil
	}

	validateErrorJSON := new(shared.ValidateErrorJSON)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&validateErrorJSON); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = UpdateSync200ApplicationJSONTypeValidateErrorJSON
		return nil
	}

	internalServerError := new(shared.InternalServerError)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&internalServerError); err == nil {
		u.InternalServerError = internalServerError
		u.Type = UpdateSync200ApplicationJSONTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdateSync200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.Sync != nil {
		return json.Marshal(u.Sync)
	}

	if u.ValidateErrorJSON != nil {
		return json.Marshal(u.ValidateErrorJSON)
	}

	if u.InternalServerError != nil {
		return json.Marshal(u.InternalServerError)
	}

	return nil, nil
}

type UpdateSyncResponse struct {
	ContentType string
	// Something went wrong
	InternalServerError *shared.InternalServerError
	StatusCode          int
	RawResponse         *http.Response
	// Ok
	UpdateSync200ApplicationJSONAnyOf *UpdateSync200ApplicationJSON
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
