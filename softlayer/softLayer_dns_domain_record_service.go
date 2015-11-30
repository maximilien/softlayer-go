package softlayer

import (
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
)

type SoftLayer_Dns_Domain_Record_Service interface {
	Service

	CreateObject(template datatypes.SoftLayer_Dns_Domain_Record) (datatypes.SoftLayer_Dns_Domain_Record, error)
	GetObject(id string) (datatypes.SoftLayer_Dns_Domain_Record, error)
	DeleteObject(id string) (datatypes.SoftLayer_Dns_Domain_Record, error)
	UpdateObject(id string) (datatypes.SoftLayer_Dns_Domain_Record, error)
}
