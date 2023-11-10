// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/de-tf-providers/terraform-provider-hightouch/v2/internal/sdk/pkg/models/shared"
	"github.com/de-tf-providers/terraform-provider-hightouch/v2/internal/sdk/pkg/utils"
	"net/http"
)

type UpdateSyncRequest struct {
	SyncUpdate shared.SyncUpdate `request:"mediaType=application/json"`
	// The sync's ID
	SyncID float64 `pathParam:"style=simple,explode=false,name=syncId"`
}

func (o *UpdateSyncRequest) GetSyncUpdate() shared.SyncUpdate {
	if o == nil {
		return shared.SyncUpdate{}
	}
	return o.SyncUpdate
}

func (o *UpdateSyncRequest) GetSyncID() float64 {
	if o == nil {
		return 0.0
	}
	return o.SyncID
}

type UpdateSyncResponseBodyType string

const (
	UpdateSyncResponseBodyTypeSync                UpdateSyncResponseBodyType = "Sync"
	UpdateSyncResponseBodyTypeValidateErrorJSON   UpdateSyncResponseBodyType = "ValidateErrorJSON"
	UpdateSyncResponseBodyTypeInternalServerError UpdateSyncResponseBodyType = "InternalServerError"
)

type UpdateSyncResponseBody struct {
	Sync                *shared.Sync
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type UpdateSyncResponseBodyType
}

func CreateUpdateSyncResponseBodySync(sync shared.Sync) UpdateSyncResponseBody {
	typ := UpdateSyncResponseBodyTypeSync

	return UpdateSyncResponseBody{
		Sync: &sync,
		Type: typ,
	}
}

func CreateUpdateSyncResponseBodyValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) UpdateSyncResponseBody {
	typ := UpdateSyncResponseBodyTypeValidateErrorJSON

	return UpdateSyncResponseBody{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateUpdateSyncResponseBodyInternalServerError(internalServerError shared.InternalServerError) UpdateSyncResponseBody {
	typ := UpdateSyncResponseBodyTypeInternalServerError

	return UpdateSyncResponseBody{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *UpdateSyncResponseBody) UnmarshalJSON(data []byte) error {

	validateErrorJSON := new(shared.ValidateErrorJSON)
	if err := utils.UnmarshalJSON(data, &validateErrorJSON, "", true, true); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = UpdateSyncResponseBodyTypeValidateErrorJSON
		return nil
	}

	sync := new(shared.Sync)
	if err := utils.UnmarshalJSON(data, &sync, "", true, true); err == nil {
		u.Sync = sync
		u.Type = UpdateSyncResponseBodyTypeSync
		return nil
	}

	internalServerError := new(shared.InternalServerError)
	if err := utils.UnmarshalJSON(data, &internalServerError, "", true, true); err == nil {
		u.InternalServerError = internalServerError
		u.Type = UpdateSyncResponseBodyTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdateSyncResponseBody) MarshalJSON() ([]byte, error) {
	if u.Sync != nil {
		return utils.MarshalJSON(u.Sync, "", true)
	}

	if u.ValidateErrorJSON != nil {
		return utils.MarshalJSON(u.ValidateErrorJSON, "", true)
	}

	if u.InternalServerError != nil {
		return utils.MarshalJSON(u.InternalServerError, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UpdateSyncResponse struct {
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
	OneOf *UpdateSyncResponseBody
}

func (o *UpdateSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSyncResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateSyncResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSyncResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSyncResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}

func (o *UpdateSyncResponse) GetOneOf() *UpdateSyncResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
