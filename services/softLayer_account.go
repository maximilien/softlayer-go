package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

type softLayer_Account_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Account_Service(client softlayer.Client) *softLayer_Account_Service {
	return &softLayer_Account_Service{
		client: client,
	}
}

func (slas *softLayer_Account_Service) GetName() string {
	return "SoftLayer_Account"
}

func (slas *softLayer_Account_Service) GetAccountStatus() (datatypes.SoftLayer_Account_Status, error) {
	path := fmt.Sprintf("%s/%s", slas.GetName(), "getAccountStatus.json")
	responseBytes, err := slas.client.DoRawHttpRequest(path, "GET", &bytes.Buffer{})
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: could not SoftLayer_Account#getAccountStatus, error message '%s'", err.Error())
		return datatypes.SoftLayer_Account_Status{}, errors.New(errorMessage)
	}

	accountStatus := datatypes.SoftLayer_Account_Status{}
	err = json.Unmarshal(responseBytes, &accountStatus)
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, err message '%s'", err.Error())
		err := errors.New(errorMessage)
		return datatypes.SoftLayer_Account_Status{}, err
	}

	return accountStatus, nil
}

func (slas *softLayer_Account_Service) GetVirtualGuests() ([]datatypes.SoftLayer_Virtual_Guest, error) {
	path := fmt.Sprintf("%s/%s", slas.GetName(), "getVirtualGuests.json")
	responseBytes, err := slas.client.DoRawHttpRequest(path, "GET", &bytes.Buffer{})
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: could not SoftLayer_Account#getVirtualGuests, error message '%s'", err.Error())
		return []datatypes.SoftLayer_Virtual_Guest{}, errors.New(errorMessage)
	}

	virtualGuests := []datatypes.SoftLayer_Virtual_Guest{}
	err = json.Unmarshal(responseBytes, &virtualGuests)
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, err message '%s'", err.Error())
		err := errors.New(errorMessage)
		return []datatypes.SoftLayer_Virtual_Guest{}, err
	}

	return virtualGuests, nil
}

func (slas *softLayer_Account_Service) GetNetworkStorage() ([]datatypes.SoftLayer_Network_Storage, error) {
	path := fmt.Sprintf("%s/%s", slas.GetName(), "getNetworkStorage.json")
	responseBytes, err := slas.client.DoRawHttpRequest(path, "GET", &bytes.Buffer{})
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: could not SoftLayer_Account#getNetworkStorage, error message '%s'", err.Error())
		return []datatypes.SoftLayer_Network_Storage{}, errors.New(errorMessage)
	}

	networkStorage := []datatypes.SoftLayer_Network_Storage{}
	err = json.Unmarshal(responseBytes, &networkStorage)
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, err message '%s'", err.Error())
		err := errors.New(errorMessage)
		return []datatypes.SoftLayer_Network_Storage{}, err
	}

	return networkStorage, nil
}

func (slas *softLayer_Account_Service) GetIscsiNetworkStorage() ([]datatypes.SoftLayer_Network_Storage, error) {
	path := fmt.Sprintf("%s/%s", slas.GetName(), "getIscsiNetworkStorage.json")

	objectMasks := []string{
		"username",
		"accountId",
		"capacityGb",
		"id",
		"billingItem.id",
		"billingItem.orderItem.order.id",
	}

	responseBytes, err := slas.client.DoRawHttpRequestWithObjectMask(path, objectMasks, "GET", &bytes.Buffer{})
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: could not SoftLayer_Account#getIscsiNetworkStorage, error message '%s'", err.Error())
		return []datatypes.SoftLayer_Network_Storage{}, errors.New(errorMessage)
	}

	networkStorage := []datatypes.SoftLayer_Network_Storage{}
	err = json.Unmarshal(responseBytes, &networkStorage)
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, err message '%s'", err.Error())
		err := errors.New(errorMessage)
		return []datatypes.SoftLayer_Network_Storage{}, err
	}

	return networkStorage, nil
}

func (slas *softLayer_Account_Service) GetVirtualDiskImages() ([]datatypes.SoftLayer_Virtual_Disk_Image, error) {
	path := fmt.Sprintf("%s/%s", slas.GetName(), "getVirtualDiskImages.json")
	responseBytes, err := slas.client.DoRawHttpRequest(path, "GET", &bytes.Buffer{})
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: could get SoftLayer_Account#getVirtualDiskImages, error message '%s'", err.Error())
		return []datatypes.SoftLayer_Virtual_Disk_Image{}, errors.New(errorMessage)
	}

	virtualDiskImages := []datatypes.SoftLayer_Virtual_Disk_Image{}
	err = json.Unmarshal(responseBytes, &virtualDiskImages)
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, err message '%s'", err.Error())
		err := errors.New(errorMessage)
		return []datatypes.SoftLayer_Virtual_Disk_Image{}, err
	}

	return virtualDiskImages, nil
}

func (slas *softLayer_Account_Service) GetSshKeys() ([]datatypes.SoftLayer_Security_Ssh_Key, error) {
	path := fmt.Sprintf("%s/%s", slas.GetName(), "getSshKeys.json")
	responseBytes, err := slas.client.DoRawHttpRequest(path, "GET", &bytes.Buffer{})
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: could not SoftLayer_Account#getSshKeys, error message '%s'", err.Error())
		return []datatypes.SoftLayer_Security_Ssh_Key{}, errors.New(errorMessage)
	}

	sshKeys := []datatypes.SoftLayer_Security_Ssh_Key{}
	err = json.Unmarshal(responseBytes, &sshKeys)
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, err message '%s'", err.Error())
		err := errors.New(errorMessage)
		return []datatypes.SoftLayer_Security_Ssh_Key{}, err
	}

	return sshKeys, nil
}

func (slas *softLayer_Account_Service) GetBlockDeviceTemplateGroups() ([]datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group, error) {
	path := fmt.Sprintf("%s/%s", slas.GetName(), "getBlockDeviceTemplateGroups.json")
	responseBytes, err := slas.client.DoRawHttpRequest(path, "GET", &bytes.Buffer{})
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: could not SoftLayer_Account#getBlockDeviceTemplateGroups, error message '%s'", err.Error())
		return []datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group{}, errors.New(errorMessage)
	}

	vgbdtGroups := []datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group{}
	err = json.Unmarshal(responseBytes, &vgbdtGroups)
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, err message '%s'", err.Error())
		err := errors.New(errorMessage)
		return []datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group{}, err
	}

	return vgbdtGroups, nil
}
