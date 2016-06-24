package data_types

type SoftLayer_Provisioning_Hook struct {
	Id         int    `json:"id"`
	AccountId  int    `json:"accountId"`
	CreateDate string `json:"createDate"`
	ModifyDate string `json:"modifyDate"`
	Name       string `json:"name"`
	TypeId     int    `json:"typeId"`
	Uri        string `json:"uri"`
}

type SoftLayer_Provisioning_Hook_Template struct {
	//Automatically assigned
	Id int `json:"id"`
	//Required
	Name   string `json:"name"`
	TypeId int    `json:"typeId"`
	Uri    string `json:"uri"`
}

type SoftLayer_Provisioning_Hook_Parameters struct {
	Parameters []SoftLayer_Provisioning_Hook_Template `json:"parameters"`
}
