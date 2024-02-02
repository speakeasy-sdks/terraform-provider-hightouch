// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/de-tf-providers/terraform-provider-hightouch/v3/internal/sdk/pkg/models/shared"
	"github.com/de-tf-providers/terraform-provider-hightouch/v3/internal/sdk/pkg/utils"
	"net/http"
)

// QueryParamOrderBy - specify the order
type QueryParamOrderBy string

const (
	QueryParamOrderByID        QueryParamOrderBy = "id"
	QueryParamOrderByName      QueryParamOrderBy = "name"
	QueryParamOrderBySlug      QueryParamOrderBy = "slug"
	QueryParamOrderByCreatedAt QueryParamOrderBy = "createdAt"
	QueryParamOrderByUpdatedAt QueryParamOrderBy = "updatedAt"
)

func (e QueryParamOrderBy) ToPointer() *QueryParamOrderBy {
	return &e
}

func (e *QueryParamOrderBy) UnmarshalJSON(data []byte) error {
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
		*e = QueryParamOrderBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamOrderBy: %v", v)
	}
}

type ListModelRequest struct {
	// limit the number of objects returned (default is 100)
	Limit *float64 `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// filter based on name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// set the offset on results (for pagination)
	Offset *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// specify the order
	OrderBy *QueryParamOrderBy `default:"id" queryParam:"style=form,explode=true,name=orderBy"`
	// filter based on slug
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (l ListModelRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListModelRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListModelRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListModelRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListModelRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListModelRequest) GetOrderBy() *QueryParamOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListModelRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

// ListModelResponseBody - Ok
type ListModelResponseBody struct {
	Data []shared.Model `json:"data"`
}

func (o *ListModelResponseBody) GetData() []shared.Model {
	if o == nil {
		return []shared.Model{}
	}
	return o.Data
}

type ListModelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
	// Ok
	Object *ListModelResponseBody
}

func (o *ListModelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListModelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListModelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListModelResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}

func (o *ListModelResponse) GetObject() *ListModelResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
