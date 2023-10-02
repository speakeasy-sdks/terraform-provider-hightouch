// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"hightouch/internal/sdk/pkg/models/shared"
	"net/http"
)

// ListModelOrderBy - specify the order
type ListModelOrderBy string

const (
	ListModelOrderByID        ListModelOrderBy = "id"
	ListModelOrderByName      ListModelOrderBy = "name"
	ListModelOrderBySlug      ListModelOrderBy = "slug"
	ListModelOrderByCreatedAt ListModelOrderBy = "createdAt"
	ListModelOrderByUpdatedAt ListModelOrderBy = "updatedAt"
)

func (e ListModelOrderBy) ToPointer() *ListModelOrderBy {
	return &e
}

func (e *ListModelOrderBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "id":
		fallthrough
	case "name":
		fallthrough
	case "slug":
		fallthrough
	case "createdAt":
		fallthrough
	case "updatedAt":
		*e = ListModelOrderBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListModelOrderBy: %v", v)
	}
}

type ListModelRequest struct {
	// limit the number of objects returned (default is 100)
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// filter based on name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// set the offset on results (for pagination)
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	// specify the order
	OrderBy *ListModelOrderBy `queryParam:"style=form,explode=true,name=orderBy"`
	// filter based on slug
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

// ListModel200ApplicationJSON - Ok
type ListModel200ApplicationJSON struct {
	Data []shared.Model `json:"data"`
}

type ListModelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Ok
	ListModel200ApplicationJSONObject *ListModel200ApplicationJSON
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
