package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type NetworkApplicationDeliveryControllerCreateOptions struct {
	Speed    int
	Version  string
	Plan     string
	IpCount  int
	Location string
}

type SoftLayer_Network_Application_Delivery_Controller_Service interface {
	Service

	CreateNetscalerVPX(createOptions *NetworkApplicationDeliveryControllerCreateOptions) (datatypes.SoftLayer_Network_Application_Delivery_Controller, error)
	CreateVirtualIpAddress(nadcId int, template datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress_Template) (bool, error)

	DeleteVirtualIpAddress(nadcId int, name string) (bool, error)
	DeleteObject(id int) (bool, error)

	EditVirtualIpAddress(nadcId int, template datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress_Template) (bool, error)

	GetObject(id int) (datatypes.SoftLayer_Network_Application_Delivery_Controller, error)
	GetVirtualIpAddress(nadcId int, vipId int) (datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress, error)
}
