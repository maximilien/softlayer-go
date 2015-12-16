package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type CreateOptions struct {
	Speed    int
	Version  string
	Plan     string
	IpCount  int
	Location string
}

type SoftLayer_Network_Application_Delivery_Controller_Service interface {
	Service

	CreateNetscalerVPX(createOptions *CreateOptions) (datatypes.SoftLayer_Network_Application_Delivery_Controller, error)

	DeleteObject(id int) (bool, error)
	GetObject(id int) (datatypes.SoftLayer_Network_Application_Delivery_Controller, error)
}
