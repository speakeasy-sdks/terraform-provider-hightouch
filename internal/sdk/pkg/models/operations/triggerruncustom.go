// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type TriggerRunCustomSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type TriggerRunCustomResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Ok
	TriggerRunCustom200ApplicationJSONAnyOf interface{}
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
