// Generated from Databricks Terraform provider schema. DO NOT EDIT.

package schema

type ResourceShareObjectPartitionValue struct {
	Name                 string `json:"name"`
	Op                   string `json:"op"`
	RecipientPropertyKey string `json:"recipient_property_key,omitempty"`
	Value                string `json:"value,omitempty"`
}

type ResourceShareObjectPartition struct {
	Value []ResourceShareObjectPartitionValue `json:"value,omitempty"`
}

type ResourceShareObject struct {
	AddedAt                  int                            `json:"added_at,omitempty"`
	AddedBy                  string                         `json:"added_by,omitempty"`
	CdfEnabled               bool                           `json:"cdf_enabled,omitempty"`
	Comment                  string                         `json:"comment,omitempty"`
	DataObjectType           string                         `json:"data_object_type"`
	HistoryDataSharingStatus string                         `json:"history_data_sharing_status,omitempty"`
	Name                     string                         `json:"name"`
	SharedAs                 string                         `json:"shared_as,omitempty"`
	StartVersion             int                            `json:"start_version,omitempty"`
	Status                   string                         `json:"status,omitempty"`
	Partition                []ResourceShareObjectPartition `json:"partition,omitempty"`
}

type ResourceShare struct {
	CreatedAt int                   `json:"created_at,omitempty"`
	CreatedBy string                `json:"created_by,omitempty"`
	Id        string                `json:"id,omitempty"`
	Name      string                `json:"name"`
	Object    []ResourceShareObject `json:"object,omitempty"`
}
