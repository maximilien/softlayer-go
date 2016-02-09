package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Dns_Domain_ResourceRecord_Service interface {
	Service

	CreateObject(template datatypes.SoftLayer_Dns_Domain_Resource_Record_Template) (datatypes.SoftLayer_Dns_Domain_Resource_Record, error)
	GetObject(recordId int) (datatypes.SoftLayer_Dns_Domain_Resource_Record, error)
	DeleteObject(recordId int) (bool, error)
	EditObject(recordId int, template datatypes.SoftLayer_Dns_Domain_Resource_Record) (bool, error)
}
