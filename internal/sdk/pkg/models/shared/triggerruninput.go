// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TriggerRunInput - The input of a trigger action to run syncs
type TriggerRunInput struct {
	// Whether to resync all the rows in the query (i.e. ignoring previously
	// synced rows).
	FullResync *bool `json:"fullResync,omitempty"`
}
