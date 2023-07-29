// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// Source - The database or warehouse where your data is stored. The starting point for
// a Hightouch data pipeline.
type Source struct {
	// The source's configuration. This specifies general metadata about sources, like connection details
	// Hightouch will use this configuration to connect to underlying source.
	//
	// The schema depends on the source type.
	//
	// Consumers should NOT make assumptions on the contents of the
	// configuration. It may change as Hightouch updates its internal code.
	Configuration map[string]interface{} `json:"configuration"`
	// The timestamp when the source was created
	CreatedAt time.Time `json:"createdAt"`
	// The source's id
	ID string `json:"id"`
	// The source's name
	Name string `json:"name"`
	// The source's slug
	Slug string `json:"slug"`
	// The source's type (e.g. snowflake or postgres).
	Type string `json:"type"`
	// The timestamp when the source was last updated
	UpdatedAt time.Time `json:"updatedAt"`
	// The id of the workspace that the source belongs to
	WorkspaceID string `json:"workspaceId"`
}
