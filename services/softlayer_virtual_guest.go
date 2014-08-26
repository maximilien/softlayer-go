package services

import (
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

	return datatypes.SoftLayer_Virtual_Guest{}, errors.New("Implement me!")
}

func (slvg *softLayerVirtualGuest) DeleteObject(template datatypes.SoftLayer_Virtual_Guest_Template) (bool, error) {
	return false, errors.New("Implement me!")
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
		errorMessage += fmt.Sprintf(errorTemplate, "datacenter.name: specifies which datacenter the instance is to be provisioned in")
	}

	if errorMessage != "" {
		err = errors.New(errorMessage)
	}

	return err
}
