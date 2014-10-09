package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service interface {
	Service

	GetObject(id int) (datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group, error)
	DeleteObject(id int) (bool, error)
	GetDatacenters(id int) ([]datatypes.SoftLayer_Location, error)
	GetSshKeys(id int) ([]datatypes.SoftLayer_Security_Ssh_Key, error)
	GetStatus(id int) (datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status, error)
}
