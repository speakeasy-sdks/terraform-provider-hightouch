// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSourceRequest struct {
	// The id of the source
	SourceID float64 `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Ok
	Source *shared.Source
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
