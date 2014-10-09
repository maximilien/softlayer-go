package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

type softLayer_Virtual_Guest_Block_Device_Template_Group_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Virtual_Guest_Block_Device_Template_Group_Service(client softlayer.Client) *softLayer_Virtual_Guest_Block_Device_Template_Group_Service {
	return &softLayer_Virtual_Guest_Block_Device_Template_Group_Service{
		client: client,
	}
}

func (slvgs *softLayer_Virtual_Guest_Block_Device_Template_Group_Service) GetName() string {
	return "SoftLayer_Virtual_Guest_Block_Device_Template_Group"
}

func (slvgbdtg *softLayer_Virtual_Guest_Block_Device_Template_Group_Service) GetObject(id int) (datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group, error) {
	response, err := slvgbdtg.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getObject.json", slvgbdtg.GetName(), id), "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group{}, err
	}

	vgbdtGroup := datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group{}
	err = json.Unmarshal(response, &vgbdtGroup)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group{}, err
	}

	return vgbdtGroup, nil
}

func (slvgbdtg *softLayer_Virtual_Guest_Block_Device_Template_Group_Service) DeleteObject(id int) (bool, error) {
	response, err := slvgbdtg.client.DoRawHttpRequest(fmt.Sprintf("%s/%d.json", slvgbdtg.GetName(), id), "DELETE", new(bytes.Buffer))

	if res := string(response[:]); res != "true" {
		return false, errors.New(fmt.Sprintf("Failed to delete instance with id '%d', got '%s' as response from the API.", id, res))
	}

	return true, err
}

func (slvgbdtg *softLayer_Virtual_Guest_Block_Device_Template_Group_Service) GetDatacenters(id int) ([]datatypes.SoftLayer_Location, error) {
	response, err := slvgbdtg.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getDatacenters.json", slvgbdtg.GetName(), id), "GET", new(bytes.Buffer))
	if err != nil {
		return []datatypes.SoftLayer_Location{}, err
	}

	locations := []datatypes.SoftLayer_Location{}
	err = json.Unmarshal(response, &locations)
	if err != nil {
		return []datatypes.SoftLayer_Location{}, err
	}

	return locations, nil
}

func (slvgbdtg *softLayer_Virtual_Guest_Block_Device_Template_Group_Service) GetSshKeys(id int) ([]datatypes.SoftLayer_Security_Ssh_Key, error) {
	response, err := slvgbdtg.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getSshKeys.json", slvgbdtg.GetName(), id), "GET", new(bytes.Buffer))
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

func (slvgbdtg *softLayer_Virtual_Guest_Block_Device_Template_Group_Service) GetStatus(id int) (datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status, error) {
	response, err := slvgbdtg.client.DoRawHttpRequest(fmt.Sprintf("%s/%d/getStatus.json", slvgbdtg.GetName(), id), "GET", new(bytes.Buffer))
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status{}, err
	}

	status := datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status{}
	err = json.Unmarshal(response, &status)
	if err != nil {
		return datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status{}, err
	}

	return status, nil
}
