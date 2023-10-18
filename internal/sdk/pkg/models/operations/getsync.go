// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSyncRequest struct {
	// The id of the sync
	SyncID float64 `pathParam:"style=simple,explode=false,name=syncId"`
}

func (o *GetSyncRequest) GetSyncID() float64 {
	if o == nil {
		return 0.0
	}
	return o.SyncID
}

type GetSyncResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	Sync *shared.Sync
}

func (o *GetSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSyncResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSyncResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSyncResponse) GetSync() *shared.Sync {
	if o == nil {
		return nil
	}
	return o.Sync
}
