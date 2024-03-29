// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/internal/utils"
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/models/shared"
	"net/http"
	"time"
)

// ListSyncQueryParamOrderBy - specify the order
type ListSyncQueryParamOrderBy string

const (
	ListSyncQueryParamOrderByID        ListSyncQueryParamOrderBy = "id"
	ListSyncQueryParamOrderByName      ListSyncQueryParamOrderBy = "name"
	ListSyncQueryParamOrderBySlug      ListSyncQueryParamOrderBy = "slug"
	ListSyncQueryParamOrderByCreatedAt ListSyncQueryParamOrderBy = "createdAt"
	ListSyncQueryParamOrderByUpdatedAt ListSyncQueryParamOrderBy = "updatedAt"
)

func (e ListSyncQueryParamOrderBy) ToPointer() *ListSyncQueryParamOrderBy {
	return &e
}

func (e *ListSyncQueryParamOrderBy) UnmarshalJSON(data []byte) error {
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
		*e = ListSyncQueryParamOrderBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListSyncQueryParamOrderBy: %v", v)
	}
}

type ListSyncRequest struct {
	// select syncs that were run after given ISO timestamp
	After *time.Time `queryParam:"style=form,explode=true,name=after"`
	// select syncs that were run before given ISO timestamp
	Before *time.Time `queryParam:"style=form,explode=true,name=before"`
	// limit the number of objects returned (default is 100)
	Limit *float64 `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// filter based on modelId
	ModelID *float64 `queryParam:"style=form,explode=true,name=modelId"`
	// set the offset on results (for pagination)
	Offset *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// specify the order
	OrderBy *ListSyncQueryParamOrderBy `default:"id" queryParam:"style=form,explode=true,name=orderBy"`
	// filter based on slug
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (l ListSyncRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListSyncRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListSyncRequest) GetAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListSyncRequest) GetBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *ListSyncRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListSyncRequest) GetModelID() *float64 {
	if o == nil {
		return nil
	}
	return o.ModelID
}

func (o *ListSyncRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListSyncRequest) GetOrderBy() *ListSyncQueryParamOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListSyncRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

// ListSyncResponseBody - Ok
type ListSyncResponseBody struct {
	Data    []shared.Sync `json:"data"`
	HasMore bool          `json:"hasMore"`
}

func (o *ListSyncResponseBody) GetData() []shared.Sync {
	if o == nil {
		return []shared.Sync{}
	}
	return o.Data
}

func (o *ListSyncResponseBody) GetHasMore() bool {
	if o == nil {
		return false
	}
	return o.HasMore
}

type ListSyncResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
	// Ok
	Object *ListSyncResponseBody
}

func (o *ListSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSyncResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSyncResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListSyncResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}

func (o *ListSyncResponse) GetObject() *ListSyncResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
