package softlayer

import (
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
)

type SoftLayer_Product_Package_Service interface {
	Service

	GetItemPrices(packageId int) ([]datatypes.SoftLayer_Item_Price, error)
	GetItemPricesBySize(packageId int, size int) ([]datatypes.SoftLayer_Item_Price, error)
}
