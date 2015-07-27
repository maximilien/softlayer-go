package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Product_Order_Service interface {
	Service

	PlaceOrder(order datatypes.SoftLayer_Product_Order) (datatypes.SoftLayer_Product_Order_Receipt, error)
	PlaceEphemeralDiskOrder(order datatypes.SoftLayer_Ephemeral_Disk_Order) (datatypes.SoftLayer_Product_Order_Receipt, error)
}
