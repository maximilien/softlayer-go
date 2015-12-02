package data_types

type SoftLayer_Dns_Domain_Template struct {
	Name		string		`json:"name"`
	ResourceRecords		[]SoftLayer_Dns_Domain_Record		`json:"resourceRecords"`
}

type SoftLayer_Dns_Domain_Template_Parameters struct {
	Parameters 	[]SoftLayer_Dns_Domain_Template		`json:"parameters"`
}

type SoftLayer_Dns_Domain struct {
	Id						int			`json:"id"`
	Name					string		`json:"name"`
	Serial					int			`json:"serial"`
	UpdateDate				string		`json:"updateDate"`
//	Account					SoftLayer_Account 		`json:"account"`
	ManagedResourceFlag		bool		`json:"managedResourceFlag"`
	ResourceRecordCount		int 		`json:"resourceRecordCount"`
	ResourceRecords			[]SoftLayer_Dns_Domain_Record		`json:"resourceRecords"`
//	Secondary				SoftLayer_Dns_Secondary						`json:"secondary"`
}