package services

import "github.com/TheWeatherCompany/softlayer-go/softlayer"

type softLayer_Dns_Domain_Record_Srv_Service struct {
	softLayer_Dns_Domain_Record_Service

	client softlayer.Client
}

func NewSoftLayer_Dns_Domain_Record_SRV_Service(client softlayer.Client) *softLayer_Dns_Domain_Record_Srv_Service {
	return &softLayer_Dns_Domain_Record_Srv_Service{
		client: client,
	}
}

func (sldr *softLayer_Dns_Domain_Record_Srv_Service) GetName() string {
	return "SoftLayer_Dns_Domain_ResourceRecord_SrvType"
}