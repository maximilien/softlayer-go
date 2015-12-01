package data_types

type SoftLayer_Dns_Domain_Record_Template_Parameters struct {
	Parameters []SoftLayer_Dns_Domain_Record_Template `json:"parameters"`
}

type SoftLayer_Dns_Domain_Record_Template struct {
	Data              string `json:"data"`
	DomainId          int    `json:"domainId"`
	Expire            int    `json:"expire"`
	Host              string `json:"host"`
	Id                int    `json:"id"`
	Minimum           int    `json:"minimum"`
	MxPriority        int    `json:"mxPriority"`
	Refresh           int    `json:"refresh"`
	ResponsiblePerson string    `json:"responsiblePerson"`
	Retry             int    `json:"retry"`
	ttl               int    `json:"ttl"`
	Type              int    `json:"type"`
}

type SoftLayer_Dns_Domain_Record struct {
	Data              string `json:"data"`
	DomainId          int    `json:"domainId"`
	Expire            int    `json:"expire"`
	Host              string `json:"host"`
	Id                int    `json:"id"`
	Minimum           int    `json:"minimum"`
	MxPriority        int    `json:"mxPriority"`
	Refresh           int    `json:"refresh"`
	ResponsiblePerson string    `json:"responsiblePerson"`
	Retry             int    `json:"retry"`
	ttl               int    `json:"ttl"`
	Type              int    `json:"type"`
}
