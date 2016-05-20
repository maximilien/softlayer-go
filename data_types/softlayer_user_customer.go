package data_types

type SoftLayer_User_Customer struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	ParentId    int    `json:"parentId"`
}
