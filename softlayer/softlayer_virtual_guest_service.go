package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Virtual_Guest_Service interface {
	Service

	GetObject(instanceId int) (datatypes.SoftLayer_Virtual_Guest, error)
	CreateObject(template datatypes.SoftLayer_Virtual_Guest_Template) (datatypes.SoftLayer_Virtual_Guest, error)
	EditObject(instanceId int, template datatypes.SoftLayer_Virtual_Guest) (bool, error)
	DeleteObject(instanceId int) (bool, error)

	IsPingable(instanceId int) (bool, error)

	GetPowerState(instanceId int) (datatypes.SoftLayer_Virtual_Guest_Power_State, error)

	GetUserData(instanceId int) ([]datatypes.SoftLayer_Virtual_Guest_Attribute, error)

	GetSshKeys(instanceId int) ([]datatypes.SoftLayer_Security_Ssh_Key, error)

	GetActiveTransaction(instanceId int) (datatypes.SoftLayer_Provisioning_Version1_Transaction, error)
	GetActiveTransactions(instanceId int) ([]datatypes.SoftLayer_Provisioning_Version1_Transaction, error)

	GetPrimaryIpAddress(instanceId int) (string, error)

	PowerCycle(instanceId int) (bool, error)
	PowerOff(instanceId int) (bool, error)
	PowerOffSoft(instanceId int) (bool, error)
	PowerOn(instanceId int) (bool, error)

	RebootDefault(instanceId int) (bool, error)
	RebootSoft(instanceId int) (bool, error)
	RebootHard(instanceId int) (bool, error)

	SetMetadata(instanceId int, metadata string) (bool, error)

	ConfigureMetadataDisk(instanceId int) (datatypes.SoftLayer_Provisioning_Version1_Transaction, error)

	AttachIscsiVolume(instanceId int, volumeId int) (string, error)
	DetachIscsiVolume(instanceId int, volumeId int) error

	AttachEphemeralDisk(instanceId int, diskSize int) error

	GetUpgradeItemPrices(instanceId int) ([]datatypes.SoftLayer_Item_Price, error)
}
