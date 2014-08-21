package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Account interface {
	Service

	GetVirtualGuests() ([]datatypes.SoftLayer_Virtual_Guest, error)
	GetNetworkStorage() ([]datatypes.SoftLayer_Network_Storage, error)
	GetVirtualDiskImages() ([]datatypes.SoftLayer_Virtual_Disk_Image, error)
}
