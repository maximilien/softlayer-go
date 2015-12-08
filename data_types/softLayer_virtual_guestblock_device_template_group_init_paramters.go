package data_types

type SoftLayer_Virtual_Guest_Block_Device_Template_GroupInitParameters struct {
	Parameters SoftLayer_Virtual_Guest_Block_Device_Template_GroupInitParameter `json:"parameters"`
}

type SoftLayer_Virtual_Guest_Block_Device_Template_GroupInitParameter struct {
	AccountId int `json:"accountId"`
}
