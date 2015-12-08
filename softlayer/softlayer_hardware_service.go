package softlayer

import (
	datatypes "github.com/maximilian/softlayer-go/data_types"
)

type SoftLayer_Hardware_Service interface {
	Service

	CreateObject(template datatypes.SoftLayer_Hardware_Template) (datatypes.SoftLayer_Hardware, error)
	GetObject(id string) (datatypes.SoftLayer_Hardware, error)
}
