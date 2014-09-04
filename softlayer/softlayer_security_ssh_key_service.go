package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Security_Ssh_Key_Service interface {
	Service

	CreateObject(template datatypes.SoftLayer_Security_Ssh_Key) (datatypes.SoftLayer_Security_Ssh_Key, error)
	DeleteObject(sshKeyId int) (bool, error)
}
