package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/TheWeatherCompany/softlayer-go/common"
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
	"github.com/TheWeatherCompany/softlayer-go/softlayer"
)

type softLayer_Provisioning_Hook_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Provisioning_Hook_Service(client softlayer.Client) *softLayer_Provisioning_Hook_Service {
	return &softLayer_Provisioning_Hook_Service{
		client: client,
	}
}

func (slphs *softLayer_Provisioning_Hook_Service) GetName() string {
	return "SoftLayer_Provisioning_Hook"
}

func (slphs *softLayer_Provisioning_Hook_Service) CreateProvisioningHook(template datatypes.SoftLayer_Provisioning_Hook_Template) (datatypes.SoftLayer_Provisioning_Hook, error) {
	parameters := datatypes.SoftLayer_Provisioning_Hook_Parameters{
		Parameters: []datatypes.SoftLayer_Provisioning_Hook_Template{{
			Id:      template.Id,
			Name:    template.Name,
			TypeId:  template.TypeId,
			Uri:     template.Uri,
		}},
	}
}