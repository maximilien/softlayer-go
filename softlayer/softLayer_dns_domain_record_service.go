package softlayer

import (
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
)

type SoftLayer_Dns_Domain_Record_Service interface {
	Service

	CreateObject(template datatypes.SoftLayer_Dns_Domain_Record_Template) (datatypes.SoftLayer_Dns_Domain_Record, error)
	CreateObjects(template []datatypes.SoftLayer_Dns_Domain_Record_Template) ([]datatypes.SoftLayer_Dns_Domain_Record, error)
	GetObject(recordId int) (datatypes.SoftLayer_Dns_Domain_Record, error)
	DeleteObject(recordId int) (bool, error)
	EditObject(recordId int, template datatypes.SoftLayer_Dns_Domain_Record) (bool, error)
}
