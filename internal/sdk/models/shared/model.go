// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/de-tf-providers/terraform-provider-hightouch/v4/internal/sdk/internal/utils"
	"time"
)

// Custom query for sources that doesn't support sql. For example, Airtable.
type Custom struct {
	Query interface{} `json:"query"`
}

func (o *Custom) GetQuery() interface{} {
	if o == nil {
		return nil
	}
	return o.Query
}

// Dbt - Query that is based on a dbt model
type Dbt struct {
	// Compiled SQL in the dbt model
	CompiledSQL string `json:"compiledSql"`
	// Name of the database containing the generated table
	Database string `json:"database"`
	// Unique ID of the model assigned by dbt (usually some combination of the schema and table name)
	DbtUniqueID string `json:"dbtUniqueId"`
	// Model id that refer to a dbt model
	ModelID string `json:"modelId"`
	// Name of the table generated by the dbt model
	Name string `json:"name"`
	// Raw SQL in the dbt model
	RawSQL string `json:"rawSql"`
	// Name of the schema containing the generated table
	Schema string `json:"schema"`
}

func (o *Dbt) GetCompiledSQL() string {
	if o == nil {
		return ""
	}
	return o.CompiledSQL
}

func (o *Dbt) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *Dbt) GetDbtUniqueID() string {
	if o == nil {
		return ""
	}
	return o.DbtUniqueID
}

func (o *Dbt) GetModelID() string {
	if o == nil {
		return ""
	}
	return o.ModelID
}

func (o *Dbt) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Dbt) GetRawSQL() string {
	if o == nil {
		return ""
	}
	return o.RawSQL
}

func (o *Dbt) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

// Raw - Standard raw SQL query
type Raw struct {
	SQL string `json:"sql"`
}

func (o *Raw) GetSQL() string {
	if o == nil {
		return ""
	}
	return o.SQL
}

// Table - Table-based query that fetches on a table instead of SQL
type Table struct {
	Name string `json:"name"`
}

func (o *Table) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// Visual query, used by audience
type Visual struct {
	Filter interface{} `json:"filter"`
	// Parent id of the schema that visual query is based on
	ParentID       string `json:"parentId"`
	PrimaryLabel   string `json:"primaryLabel"`
	SecondaryLabel string `json:"secondaryLabel"`
}

func (o *Visual) GetFilter() interface{} {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *Visual) GetParentID() string {
	if o == nil {
		return ""
	}
	return o.ParentID
}

func (o *Visual) GetPrimaryLabel() string {
	if o == nil {
		return ""
	}
	return o.PrimaryLabel
}

func (o *Visual) GetSecondaryLabel() string {
	if o == nil {
		return ""
	}
	return o.SecondaryLabel
}

// Model - The SQL query that pulls data from your source to send to your destination.
// We send your SQL query directly to your source so any SQL that is valid for your source (including functions) is valid in Hightouch.
type Model struct {
	// The timestamp when model was created
	CreatedAt time.Time `json:"createdAt"`
	// Custom query for sources that doesn't support sql. For example, Airtable.
	Custom *Custom `json:"custom,omitempty"`
	// Query that is based on a dbt model
	Dbt *Dbt `json:"dbt,omitempty"`
	// The id of the model
	ID string `json:"id"`
	// If is_schema is true, the model is just used to build other models.
	// Either as part of visual querying, or as the root of a visual query.
	IsSchema bool `json:"isSchema"`
	// The name of the model
	Name string `json:"name"`
	// The primary key will be null if the query doesn't get directly synced (e.g. a relationship table for visual querying)
	PrimaryKey string `json:"primaryKey"`
	// The type of the query. Available options: custom, raw_sql, tabel, dbt and visual.
	QueryType string `json:"queryType"`
	// Standard raw SQL query
	Raw *Raw `json:"raw,omitempty"`
	// The slug of the model
	Slug string `json:"slug"`
	// The id of the source that model is connected to
	SourceID string `json:"sourceId"`
	// The list of id of syncs that uses this model
	Syncs []string `json:"syncs"`
	// Table-based query that fetches on a table instead of SQL
	Table *Table `json:"table,omitempty"`
	// The tags of the model
	Tags map[string]string `json:"tags"`
	// The timestamp when model was lastly updated
	UpdatedAt time.Time `json:"updatedAt"`
	// Visual query, used by audience
	Visual *Visual `json:"visual,omitempty"`
	// The id of the workspace where the model belongs to
	WorkspaceID string `json:"workspaceId"`
}

func (m Model) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *Model) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Model) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Model) GetCustom() *Custom {
	if o == nil {
		return nil
	}
	return o.Custom
}

func (o *Model) GetDbt() *Dbt {
	if o == nil {
		return nil
	}
	return o.Dbt
}

func (o *Model) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Model) GetIsSchema() bool {
	if o == nil {
		return false
	}
	return o.IsSchema
}

func (o *Model) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Model) GetPrimaryKey() string {
	if o == nil {
		return ""
	}
	return o.PrimaryKey
}

func (o *Model) GetQueryType() string {
	if o == nil {
		return ""
	}
	return o.QueryType
}

func (o *Model) GetRaw() *Raw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *Model) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *Model) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *Model) GetSyncs() []string {
	if o == nil {
		return []string{}
	}
	return o.Syncs
}

func (o *Model) GetTable() *Table {
	if o == nil {
		return nil
	}
	return o.Table
}

func (o *Model) GetTags() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Tags
}

func (o *Model) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *Model) GetVisual() *Visual {
	if o == nil {
		return nil
	}
	return o.Visual
}

func (o *Model) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
