package data_types

type SoftLayer_Product_Item_Price struct {
	Id              uint32     `json:"id"`
	LocationGroupId uint32     `json:"locationGroupId"`
	Categories      []Category `json:"categories,omitempty"`
	Item            *Item      `json:"item,omitempty"`
}

type Item struct {
	Id          uint32 `json:"id"`
	Description string `json:"description"`
	Capacity    string `json:"capacity"`
}

type Category struct {
	Id           uint32 `json:"id"`
	CategoryCode string `json:"categoryCode"`
}
