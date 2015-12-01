package softlayer

import (
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
)

type SoftLayer_Dns_Domain_Record_Service interface {
	Service

	CreateObject(template datatypes.SoftLayer_Dns_Domain_Record_Template) (datatypes.SoftLayer_Dns_Domain_Record, error)
	CreateObjects(templates []datatypes.SoftLayer_Dns_Domain_Record_Template) ([]datatypes.SoftLayer_Dns_Domain_Record, error)
	GetObject(recordId int) (datatypes.SoftLayer_Dns_Domain_Record, error)
	DeleteObject(recordId int) (bool, error)
	DeleteObjects(recordIds []int) (bool, error)
	GetDomain() (datatypes.SoftLayer_Dns_Domain, error)
	UpdateObject(recordId int) (bool, error)
}
