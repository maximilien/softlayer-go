package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Network_Storage_Service interface {
	Service

	CreateIscsiVolume(size int, location string) (datatypes.SoftLayer_Network_Storage, error)
	DeleteIscsiVolume(volumeId int, immediateCancellationFlag bool) error
	GetIscsiVolume(volumeId int) (datatypes.SoftLayer_Network_Storage, error)

}
