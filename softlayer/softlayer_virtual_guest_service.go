package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Virtual_Guest_Service interface {
	Service

	CreateObject(template datatypes.SoftLayer_Virtual_Guest_Template) (datatypes.SoftLayer_Virtual_Guest, error)
	DeleteObject(instanceId int) (bool, error)
}
