package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Virtual_Guest_Service interface {
	Service

	AttachEphemeralDisk(instanceId int, diskSize int) error

	ConfigureMetadataDisk(instanceId int) (datatypes.SoftLayer_Provisioning_Version1_Transaction, error)
	CreateObject(template datatypes.SoftLayer_Virtual_Guest_Template) (datatypes.SoftLayer_Virtual_Guest, error)

	DeleteObject(instanceId int) (bool, error)

	EditObject(instanceId int, template datatypes.SoftLayer_Virtual_Guest) (bool, error)

	IsPingable(instanceId int) (bool, error)

	GetActiveTransaction(instanceId int) (datatypes.SoftLayer_Provisioning_Version1_Transaction, error)
	GetActiveTransactions(instanceId int) ([]datatypes.SoftLayer_Provisioning_Version1_Transaction, error)
	GetObject(instanceId int) (datatypes.SoftLayer_Virtual_Guest, error)
	GetPrimaryIpAddress(instanceId int) (string, error)
	GetPowerState(instanceId int) (datatypes.SoftLayer_Virtual_Guest_Power_State, error)
	GetSshKeys(instanceId int) ([]datatypes.SoftLayer_Security_Ssh_Key, error)
	GetTagReferences(instanceId int) ([]datatypes.SoftLayer_Tag_Reference, error)
	GetUpgradeItemPrices(instanceId int) ([]datatypes.SoftLayer_Item_Price, error)
	GetUserData(instanceId int) ([]datatypes.SoftLayer_Virtual_Guest_Attribute, error)

	PowerCycle(instanceId int) (bool, error)
	PowerOff(instanceId int) (bool, error)
	PowerOffSoft(instanceId int) (bool, error)
	PowerOn(instanceId int) (bool, error)

	RebootDefault(instanceId int) (bool, error)
	RebootSoft(instanceId int) (bool, error)
	RebootHard(instanceId int) (bool, error)

	SetMetadata(instanceId int, metadata string) (bool, error)
	SetTags(instanceId int, tags []string) (bool, error)
}
