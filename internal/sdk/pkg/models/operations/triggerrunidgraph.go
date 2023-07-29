// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type TriggerRunIDGraphSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type TriggerRunIDGraphRequest struct {
	TriggerRunIDGraphInput *shared.TriggerRunIDGraphInput `request:"mediaType=application/json"`
	GraphID                string                         `pathParam:"style=simple,explode=false,name=graphId"`
}

type TriggerRunIDGraphResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Ok
	TriggerRunIDGraphOutput *shared.TriggerRunIDGraphOutput
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
