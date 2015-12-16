package data_types

type SoftLayer_Network_Application_Delivery_Controller struct {
	Id                  int                  `json:"id"`
	Name                string               `json:"name"`
	TypeId              int                  `json:"typeId"`
	ModifyDate          string               `json:"modifyDate"`
	Description         string               `json:"description"`
	ManagedResourceFlag bool                 `json:"managedResourceFlag"`
	ManagementIpAddress string               `json:"managementIpAddress"`
	PrimaryIpAddress    string               `json:"primaryIpAddress"`
	Password            []SoftLayer_Password `json:"password"`

	Type	 SoftLayer_Network_Application_Delivery_Controller_Type 	`json:"type"`
}

type SoftLayer_Network_Application_Delivery_Controller_Type struct {
	KeyName string `json:"keyName"`
	Name    string `json:"name"`
}
