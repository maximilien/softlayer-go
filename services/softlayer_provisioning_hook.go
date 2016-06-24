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
	err := slphs.checkCreateObjectRequiredValues(template)
	if err != nil {
		return datatypes.SoftLayer_Provisioning_Hook{}, err
	}

	parameters := datatypes.SoftLayer_Provisioning_Hook_Parameters{
		Parameters: []datatypes.SoftLayer_Provisioning_Hook_Template{
			template,
		},
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return datatypes.SoftLayer_Provisioning_Hook{}, fmt.Errorf("Unable to create JSON: %s", err)
	}

	response, errorCode, err := slphs.client.GetHttpClient().DoRawHttpRequest(fmt.Sprintf("%s/createObject.json", slphs.GetName()), "POST", bytes.NewBuffer(requestBody))

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not create SoftLayer_Provisioning_Hook, HTTP error code: '%d", errorCode)
		return datatypes.SoftLayer_Provisioning_Hook{}, errors.New(errorMessage)
	}

	err = slphs.client.GetHttpClient().CheckForHttpResponseErrors(response)
	if err != nil {
		return datatypes.SoftLayer_Provisioning_Hook{}, err
	}

	provisioningHook := datatypes.SoftLayer_Provisioning_Hook{}
	err = json.Unmarshal(response, &provisioningHook)
	if err != nil {
		return datatypes.SoftLayer_Provisioning_Hook{}, err
	}

	return provisioningHook, nil
}

func (slvgs *softLayer_Provisioning_Hook_Service) checkCreateObjectRequiredValues(template datatypes.SoftLayer_Provisioning_Hook_Template) error {
	var err error
	errorMessage, errorTemplate := "", "* %s is required and cannot be empty\n"

	if template.Name == "" {
		errorMessage += fmt.Sprintf(errorTemplate, "Name for the Post-Install Script")
	}

	if template.TypeId == 0 {
		errorMessage += fmt.Sprintf(errorTemplate, "TypeId for the Post-Install Script")
	}

	if template.Uri <= "" {
		errorMessage += fmt.Sprintf(errorTemplate, "Uri for the Post-Install Script")
	}

	if errorMessage != "" {
		err = errors.New(errorMessage)
	}

	return err
}

func (slphs *softLayer_Provisioning_Hook_Service) GetObject(id int) (datatypes.SoftLayer_Provisioning_Hook, error) {
	response, errorCode, err := slphs.client.GetHttpClient().DoRawHttpRequest(fmt.Sprintf("%s/%d/getObject.json", slphs.GetName(), id), "GET", new(bytes.Buffer))

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not retrieve SoftLayer_Provisioning_Hook, HTTP error code: '%d'", errorCode)
		return datatypes.SoftLayer_Provisioning_Hook{}, errors.New(errorMessage)
	}

	if err != nil {
		return datatypes.SoftLayer_Provisioning_Hook{}, err
	}

	provisioningHook := datatypes.SoftLayer_Provisioning_Hook{}
	err = json.Unmarshal(response, &provisioningHook)
	if err != nil {
		return datatypes.SoftLayer_Provisioning_Hook{}, err
	}

	return provisioningHook, nil
}

func (slphs *softLayer_Provisioning_Hook_Service) DeleteObject(id int) (bool, error) {
	response, errorCode, err := slphs.client.GetHttpClient().DoRawHttpRequest(fmt.Sprintf("%s/%d.json", slphs.GetName(), id), "DELETE", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to delete Provisioning Hook with id '%d', got '%s' as a response from the SLAPI.", id, res))
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not remove SoftLayer_Provisioning_Hook with Id: %d, HTTP error code: '%d'", id, errorCode)
		return false, errors.New(errorMessage)
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
