// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ModelCustom - Custom query for sources that doesn't support sql. For example, Airtable.
type ModelCustom struct {
	Query interface{} `json:"query"`
}

// ModelDbt - Query that is based on a dbt model
type ModelDbt struct {
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

// ModelRaw - Standard raw SQL query
type ModelRaw struct {
	SQL string `json:"sql"`
}

// ModelTable - Table-based query that fetches on a table instead of SQL
type ModelTable struct {
	Name string `json:"name"`
}

// ModelVisual - Visual query, used by audience
type ModelVisual struct {
	Filter interface{} `json:"filter"`
	// Parent id of the schema that visual query is based on
	ParentID       string `json:"parentId"`
	PrimaryLabel   string `json:"primaryLabel"`
	SecondaryLabel string `json:"secondaryLabel"`
}

type InternalServerError string

const (
	InternalServerErrorInternalServerError InternalServerError = "Internal Server Error"
)

func (e InternalServerError) ToPointer() *InternalServerError {
	return &e
}

func (e *InternalServerError) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Internal Server Error":
		*e = InternalServerError(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InternalServerError: %v", v)
	}
}
