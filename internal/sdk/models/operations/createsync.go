// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/internal/utils"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/models/shared"
	"net/http"
)

type CreateSyncResponseBodyType string

const (
	CreateSyncResponseBodyTypeSync                CreateSyncResponseBodyType = "Sync"
	CreateSyncResponseBodyTypeValidateErrorJSON   CreateSyncResponseBodyType = "ValidateErrorJSON"
	CreateSyncResponseBodyTypeInternalServerError CreateSyncResponseBodyType = "InternalServerError"
)

// CreateSyncResponseBody - Ok
type CreateSyncResponseBody struct {
	Sync                *shared.Sync
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type CreateSyncResponseBodyType
}

func CreateCreateSyncResponseBodySync(sync shared.Sync) CreateSyncResponseBody {
	typ := CreateSyncResponseBodyTypeSync

	return CreateSyncResponseBody{
		Sync: &sync,
		Type: typ,
	}
}

func CreateCreateSyncResponseBodyValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) CreateSyncResponseBody {
	typ := CreateSyncResponseBodyTypeValidateErrorJSON

	return CreateSyncResponseBody{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateCreateSyncResponseBodyInternalServerError(internalServerError shared.InternalServerError) CreateSyncResponseBody {
	typ := CreateSyncResponseBodyTypeInternalServerError

	return CreateSyncResponseBody{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *CreateSyncResponseBody) UnmarshalJSON(data []byte) error {

	validateErrorJSON := shared.ValidateErrorJSON{}
	if err := utils.UnmarshalJSON(data, &validateErrorJSON, "", true, true); err == nil {
		u.ValidateErrorJSON = &validateErrorJSON
		u.Type = CreateSyncResponseBodyTypeValidateErrorJSON
		return nil
	}

	sync := shared.Sync{}
	if err := utils.UnmarshalJSON(data, &sync, "", true, true); err == nil {
		u.Sync = &sync
		u.Type = CreateSyncResponseBodyTypeSync
		return nil
	}

	internalServerError := shared.InternalServerError("")
	if err := utils.UnmarshalJSON(data, &internalServerError, "", true, true); err == nil {
		u.InternalServerError = &internalServerError
		u.Type = CreateSyncResponseBodyTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateSyncResponseBody) MarshalJSON() ([]byte, error) {
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

type CreateSyncResponse struct {
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
	OneOf *CreateSyncResponseBody
}

func (o *CreateSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSyncResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *CreateSyncResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSyncResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSyncResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}

func (o *CreateSyncResponse) GetOneOf() *CreateSyncResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
