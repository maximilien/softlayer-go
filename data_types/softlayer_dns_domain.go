package data_types

type SoftLayer_Dns_Domain_Template struct {
	Name            string                                 `json:"name"`
	ResourceRecords []SoftLayer_Dns_Domain_Resource_Record `json:"resourceRecords"`
}

type SoftLayer_Dns_Domain_Template_Parameters struct {
	Parameters []SoftLayer_Dns_Domain_Template `json:"parameters"`
}

type SoftLayer_Dns_Domain struct {
	Id                  uint32                                 `json:"id"`
	Name                string                                 `json:"name"`
	Serial              uint32                                 `json:"serial"`
	UpdateDate          string                                 `json:"updateDate"`
	ManagedResourceFlag bool                                   `json:"managedResourceFlag"`
	ResourceRecordCount int                                    `json:"resourceRecordCount"`
	ResourceRecords     []SoftLayer_Dns_Domain_Resource_Record `json:"resourceRecords"`
}
