package data_types

type SoftLayer_Dns_Domain_Record_Template_Parameters struct {
	Parameters []SoftLayer_Hardware_Template `json:"parameters"`
}

type SoftLayer_Dns_Domain_Record_Template struct {
	Host                     string `json:"hostname"`
}

type SoftLayer_Dns_Domain_Record struct {
	Domain                string     `json:"domain"`
	Host                     string `json:"hostname"`
}
