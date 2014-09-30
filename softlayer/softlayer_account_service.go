package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Account_Service interface {
	Service

	GetAccountStatus() (datatypes.SoftLayer_Account_Status, error)
	GetVirtualGuests() ([]datatypes.SoftLayer_Virtual_Guest, error)
	GetNetworkStorage() ([]datatypes.SoftLayer_Network_Storage, error)
	GetIscsiNetworkStorage() ([]datatypes.SoftLayer_Network_Storage, error)
	GetVirtualDiskImages() ([]datatypes.SoftLayer_Virtual_Disk_Image, error)
	GetSshKeys() ([]datatypes.SoftLayer_Security_Ssh_Key, error)
	GetBlockDeviceTemplateGroups() ([]datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group, error)
}
