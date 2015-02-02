package data_types

type SoftLayer_Product_Order_Receipt struct {
	OrderId int `json:"orderId"`
}

type SoftLayer_Product_Order_Parameters struct {
	Parameters []SoftLayer_Product_Order `json:"parameters"`
}

type SoftLayer_Product_Order struct {
	ComplexType   string                 `json:"complexType"`
	Location      string                 `json:"location,omitempty"`
	PackageId     int                    `json:"packageId"`
	Prices        []SoftLayer_Item_Price `json:"prices,omitempty"`
	VirtualGuests []VirtualGuest         `json:"virtualGuests,omitempty"`
	Properties    []Property             `json:"properties,omitempty"`
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type VirtualGuest struct {
	Id int `json:"id"`
}
