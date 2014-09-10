package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

type softLayer_Virtual_Guest_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Virtual_Guest_Service(client softlayer.Client) *softLayer_Virtual_Guest_Service {
	return &softLayer_Virtual_Guest_Service{
		client: client,
	}
}

func (slvgs *softLayer_Virtual_Guest_Service) GetName() string {
	return "SoftLayer_Virtual_Guest"
}

func (slvgs *softLayer_Virtual_Guest_Service) CreateObject(template datatypes.SoftLayer_Virtual_Guest_Template) (datatypes.SoftLayer_Virtual_Guest, error) {
	err := slvgs.checkCreateObjectRequiredValues(template)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	parameters := datatypes.SoftLayer_Virtual_Guest_Template_Parameters{
		Parameters: []datatypes.SoftLayer_Virtual_Guest_Template{
			template,
		},
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s.json", slvgs.GetName()), "POST", bytes.NewBuffer(requestBody))
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	err = slvgs.client.CheckForHttpResponseErrors(response)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	softLayer_Virtual_Guest := datatypes.SoftLayer_Virtual_Guest{}
	err = json.Unmarshal(response, &softLayer_Virtual_Guest)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest{}, err
	}

	return softLayer_Virtual_Guest, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) EditObject(instanceId int, template datatypes.SoftLayer_Virtual_Guest) (bool, error) {
	parameters := datatypes.SoftLayer_Virtual_Guest_Parameters{
		Parameters: []datatypes.SoftLayer_Virtual_Guest{template},
	}

	requestBody, err := json.Marshal(parameters)
	if err != nil {
		return false, err
	}

	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/editObject.json", slvgs.GetName(), instanceId), "POST", bytes.NewBuffer(requestBody))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to edit virtual guest with id: %d, got '%s' as response from the API.", instanceId, res))
	}

	return true, err
}

func (slvgs *softLayer_Virtual_Guest_Service) DeleteObject(instanceId int) (bool, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d.json", slvgs.GetName(), instanceId), "DELETE", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to destroy and instance with id '%d', got '%s' as response from the API.", instanceId, res))
	}

	return true, err
}

func (slvgs *softLayer_Virtual_Guest_Service) GetPowerState(instanceId int) (datatypes.SoftLayer_Virtual_Guest_Power_State, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getPowerState.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest_Power_State{}, err
	}

	vgPowerState := datatypes.SoftLayer_Virtual_Guest_Power_State{}
	err = json.Unmarshal(response, &vgPowerState)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest_Power_State{}, err
	}

	return vgPowerState, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) GetActiveTransaction(instanceId int) (datatypes.SoftLayer_Provisioning_Version1_Transaction, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getActiveTransaction.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Provisioning_Version1_Transaction{}, err
	}

	activeTransaction := datatypes.SoftLayer_Provisioning_Version1_Transaction{}
	err = json.Unmarshal(response, &activeTransaction)
	if err != nil {
		return datatypes.SoftLayer_Provisioning_Version1_Transaction{}, err
	}

	return activeTransaction, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) GetActiveTransactions(instanceId int) ([]datatypes.SoftLayer_Provisioning_Version1_Transaction, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getActiveTransactions.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))
	if err != nil {
		return []datatypes.SoftLayer_Provisioning_Version1_Transaction{}, err
	}

	activeTransactions := []datatypes.SoftLayer_Provisioning_Version1_Transaction{}
	err = json.Unmarshal(response, &activeTransactions)
	if err != nil {
		return []datatypes.SoftLayer_Provisioning_Version1_Transaction{}, err
	}

	return activeTransactions, nil
}

func (slvgs *softLayer_Virtual_Guest_Service) GetSshKeys(instanceId int) ([]datatypes.SoftLayer_Security_Ssh_Key, error) {
	response, err := slvgs.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getSshKeys.json", slvgs.GetName(), instanceId), "GET", new(bytes.Buffer))
	if err != nil {
		return []datatypes.SoftLayer_Security_Ssh_Key{}, err
	}

	sshKeys := []datatypes.SoftLayer_Security_Ssh_Key{}
	err = json.Unmarshal(response, &sshKeys)
	if err != nil {
		return []datatypes.SoftLayer_Security_Ssh_Key{}, err
	}

	return sshKeys, nil
}

//Private methods

func (slvgs *softLayer_Virtual_Guest_Service) checkCreateObjectRequiredValues(template datatypes.SoftLayer_Virtual_Guest_Template) error {
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
