package data_types

type SoftLayer_Network_LoadBalancer_VirtualIpAddress_Array []SoftLayer_Network_LoadBalancer_VirtualIpAddress

type SoftLayer_Network_LoadBalancer_VirtualIpAddress struct {
	ConnectionLimit             int    `json:"connectionLimit"`
	CustomManagedFlag           bool   `json:"customManagedFlag"`
	Id                          int    `json:"id"`
	LoadBalancingMethod         string `json:"loadBalancingMethod"`
	LoadBalancingMethodFullName string `json:"loadBalancingMethodFullName"`
	ModifyDate                  string `json:"modifyDate"`
	Name                        string `json:"name"`
	Notes                       string `json:"notes"`
	SecurityCertificateId       int    `json:"securityCertificateId"`
	SourcePort                  int    `json:"sourcePort"`
	Type                        string `json:"type"`
	VirtualIpAddress            string `json:"virtualIpAddress"`
}

type SoftLayer_Network_LoadBalancer_VirtualIpAddress_Template_Parameters struct {
	LoadBalancer SoftLayer_Network_LoadBalancer_VirtualIpAddress_Template `json:"loadBalancer"`
}

type SoftLayer_Network_LoadBalancer_VirtualIpAddress_Template struct {
	Id                          int    `json:"id"`
	ConnectionLimit             int    `json:"connectionLimit"`
	CustomManagedFlag           bool   `json:"customManagedFlag"`
	LoadBalancingMethod         string `json:"loadBalancingMethod"`
	LoadBalancingMethodFullName string `json:"loadBalancingMethodFullName"`
	Name                        string `json:"name"`
	Notes                       string `json:"notes"`
	SecurityCertificateId       int    `json:"securityCertificateId"`
	SourcePort                  int    `json:"sourcePort"`
	Type                        string `json:"type"`
	VirtualIpAddress            string `json:"virtualIpAddress"`
}
