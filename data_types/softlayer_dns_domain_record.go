package data_types

type SoftLayer_Dns_Domain_Resource_Record_Template_Parameters struct {
	Parameters []SoftLayer_Dns_Domain_Resource_Record_Template `json:"parameters"`
}

type SoftLayer_Dns_Domain_Resource_Record_Template struct {
	Data              string `json:"data"`
	DomainId          uint32 `json:"domainId"`
	Expire            int    `json:"expire"`
	Host              string `json:"host"`
	Id                uint32 `json:"id"`
	Minimum           int    `json:"minimum"`
	MxPriority        int    `json:"mxPriority"`
	Refresh           int    `json:"refresh"`
	ResponsiblePerson string `json:"responsiblePerson"`
	Retry             int    `json:"retry"`
	Ttl               int    `json:"ttl"`
	Type              string `json:"type"`
	Service           string `json:"service,omitempty"`
	Protocol          string `json:"protocol,omitempty"`
	Priority          int    `json:"priority,omitempty"`
	Port              int    `json:"port,omitempty"`
	Weight            int    `json:"weight,omitempty"`
}

type SoftLayer_Dns_Domain_Resource_Record_Parameters struct {
	Parameters []SoftLayer_Dns_Domain_Resource_Record `json:"parameters"`
}

type SoftLayer_Dns_Domain_Resource_Record struct {
	Data              string `json:"data"`
	DomainId          uint32 `json:"domainId"`
	Expire            int    `json:"expire"`
	Host              string `json:"host"`
	Id                uint32 `json:"id"`
	Minimum           int    `json:"minimum"`
	MxPriority        int    `json:"mxPriority"`
	Refresh           int    `json:"refresh"`
	ResponsiblePerson string `json:"responsiblePerson"`
	Retry             int    `json:"retry"`
	Ttl               int    `json:"ttl"`
	Type              string `json:"type"`
	Service           string `json:"service,omitempty"`
	Protocol          string `json:"protocol,omitempty"`
	Priority          int    `json:"priority,omitempty"`
	Port              int    `json:"port,omitempty"`
	Weight            int    `json:"weight,omitempty"`
}
