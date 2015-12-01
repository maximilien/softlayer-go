package data_types

import "time"

type SoftLayer_Dns_Domain_Template struct {
	Name		string		`json:"name,omitempty"`
	Serial		int			`json:"serial,omitempty"`

	//	ResourceRecords		[]Softlayer_Dns_Domain_Record		`json:"resourceRecords,omitempty"`
}

type SoftLayer_Dns_Domain_Template_Parameters struct {
	Parameters 	[]SoftLayer_Dns_Domain_Template		`json:"parameters"`
}

type SoftLayer_Dns_Domain struct {
	Id						int			`json:"id,omitempty"`
	Name					string		`json:"name,omitempty"`
	Serial					int			`json:"serial,omitempty"`
	UpdateDate				*time.Time	`json:"updateDate,omitempty"`
//	Account					SoftLayer_Account 		`json:"account,omitempty"`
	ManagedResourceFlag		bool		`json:"managedResourceFlag,omitempty"`
	ResourceRecordCount		int 		`json:"resourceRecordCount,omitempty"`
//	ResourceRecords			[]Softlayer_Dns_Domain_Record		`json:"resourceRecords,omitempty"`
//	Secondary				SoftLayer_Dns_Secondary						`json:"secondary,omitempty"`
}