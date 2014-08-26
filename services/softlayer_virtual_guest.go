package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

type softLayerVirtualGuest struct {
	client softlayer.Client
}

func NewSoftLayer_Virtual_Guest(client softlayer.Client) *softLayerVirtualGuest {
	return &softLayerVirtualGuest{
		client: client,
	}
}

func (slvg *softLayerVirtualGuest) GetName() string {
	return "SoftLayer_Virtual_Guest"
}

func (slvg *softLayerVirtualGuest) CreateObject(template datatypes.SoftLayer_Virtual_Guest_Template) (datatypes.SoftLayer_Virtual_Guest, error) {
	err := slvg.checkCreateObjectRequiredValues(template)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	requestBody, err := json.Marshal(template)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	data, err := slvg.client.DoRawHttpRequest("SoftLayer_Virtual_Guest/createObject", "POST", bytes.NewBuffer(requestBody))
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	var decodedResponse map[string]interface{}
	err = json.Unmarshal(data, decodedResponse)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	if err := slvg.client.HasErrors(decodedResponse); err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	softLayerVirtualGuest := datatypes.SoftLayer_Virtual_Guest{}
	err = json.Unmarshal(data, softLayerVirtualGuest)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	return softLayerVirtualGuest, errors.New("Implement me!")
}

func (slvg *softLayerVirtualGuest) DeleteObject(instanceId int) (bool, error) {
	response, err := slvg.client.DoRawHttpRequest(fmt.Sprintf("SoftLayer_Virtual_Guest/%d.json", instanceId), "DELETE", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to destroy and instance wit id '%d', got '%s' as response from the API.", instanceId, res))
	}

	return true, err
}

//Private methods

func (slvg *softLayerVirtualGuest) checkCreateObjectRequiredValues(template datatypes.SoftLayer_Virtual_Guest_Template) error {
	var err error
	errorMessage, errorTemplate := "", "* %s is required and cannot be empty\n"

	if template.Hostname == "" {
		errorMessage += fmt.Sprintf(errorTemplate, "Hostname for the computing instance")
	}

	if template.Domain == "" {
		errorMessage += fmt.Sprintf(errorTemplate, "Domain for the computing instance")
	}

	if template.StartCpus <= 0 {
		errorMessage += fmt.Sprintf(errorTemplate, "StartCpus: the number of CPU cores to allocate")
	}

	if template.MaxMemory <= 0 {
		errorMessage += fmt.Sprintf(errorTemplate, "MaxMemory: the amount of memory to allocate in megabytes")
	}

	if template.Datacenter.Name == "" {
		errorMessage += fmt.Sprintf(errorTemplate, "Datacenter.Name: specifies which datacenter the instance is to be provisioned in")
	}

	if errorMessage != "" {
		err = errors.New(errorMessage)
	}

	return err
}
