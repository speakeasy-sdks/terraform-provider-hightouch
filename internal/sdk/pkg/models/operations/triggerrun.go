// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type TriggerRunSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type TriggerRunRequest struct {
	TriggerRunInput *shared.TriggerRunInput `request:"mediaType=application/json"`
	// The id of the sync to trigger a run
	SyncID string `pathParam:"style=simple,explode=false,name=syncId"`
}

type TriggerRunResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Ok
	TriggerRunOutput *shared.TriggerRunOutput
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
