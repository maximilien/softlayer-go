package data_types

type UserMetadata string

type SoftLayer_SetUserMetadata_Parameters struct {
	Parameters []UserMetadata `json:"parameters"`
}