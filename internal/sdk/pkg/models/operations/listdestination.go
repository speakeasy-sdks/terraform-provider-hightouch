// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/pkg/models/shared"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/pkg/utils"
	"net/http"
)

// OrderBy - Order the returned destinations
type OrderBy string

const (
	OrderByID        OrderBy = "id"
	OrderByName      OrderBy = "name"
	OrderBySlug      OrderBy = "slug"
	OrderByCreatedAt OrderBy = "createdAt"
	OrderByUpdatedAt OrderBy = "updatedAt"
)

func (e OrderBy) ToPointer() *OrderBy {
	return &e
}

func (e *OrderBy) UnmarshalJSON(data []byte) error {
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
		*e = OrderBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OrderBy: %v", v)
	}
}

type ListDestinationRequest struct {
	// limit the number of objects returned (default is 100)
	Limit *float64 `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Filter based on the destination's name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// set the offset on results (for pagination)
	Offset *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// Order the returned destinations
	OrderBy *OrderBy `default:"id" queryParam:"style=form,explode=true,name=orderBy"`
	// Filter based on destination's slug
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (l ListDestinationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListDestinationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListDestinationRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListDestinationRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListDestinationRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListDestinationRequest) GetOrderBy() *OrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListDestinationRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

// ListDestinationResponseBody - Ok
type ListDestinationResponseBody struct {
	Data []shared.Destination `json:"data"`
}

func (o *ListDestinationResponseBody) GetData() []shared.Destination {
	if o == nil {
		return []shared.Destination{}
	}
	return o.Data
}

type ListDestinationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
	// Ok
	Object *ListDestinationResponseBody
}

func (o *ListDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListDestinationResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}

func (o *ListDestinationResponse) GetObject() *ListDestinationResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
