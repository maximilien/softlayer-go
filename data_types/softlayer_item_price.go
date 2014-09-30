package data_types

type SoftLayer_Item_Price struct {
	Id   int   `json:"id"`
	Item *Item `json:"item"`
}

type Item struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Capacity    string `json:"capacity"`
}
