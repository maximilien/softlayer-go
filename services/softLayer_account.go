package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

const NAME = "SoftLayer_Account"

type softLayer_Account struct {
	client softlayer.Client
}

func NewSoftLayer_Account(client softlayer.Client) *softLayer_Account {
	return &softLayer_Account{
		client: client,
	}
}

func (sla *softLayer_Account) GetName() string {
	return NAME
}

func (sla *softLayer_Account) GetAccountStatus() (datatypes.SoftLayer_Account_Status, error) {
	path := fmt.Sprintf("%s/%s", sla.GetName(), "getAccountStatus.json")
	responseBytes, err := sla.client.DoRawHttpRequest(path, "GET", &bytes.Buffer{})
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

func (sla *softLayer_Account) GetVirtualGuests() ([]datatypes.SoftLayer_Virtual_Guest, error) {
	path := fmt.Sprintf("%s/%s", sla.GetName(), "getVirtualGuests.json")
	responseBytes, err := sla.client.DoRawHttpRequest(path, "GET", &bytes.Buffer{})
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

func (sla *softLayer_Account) GetNetworkStorage() ([]datatypes.SoftLayer_Network_Storage, error) {
	path := fmt.Sprintf("%s/%s", sla.GetName(), "getNetworkStorage.json")
	responseBytes, err := sla.client.DoRawHttpRequest(path, "GET", &bytes.Buffer{})
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

func (sla *softLayer_Account) GetVirtualDiskImages() ([]datatypes.SoftLayer_Virtual_Disk_Image, error) {
	path := fmt.Sprintf("%s/%s", sla.GetName(), "getVirtualDiskImages.json")
	responseBytes, err := sla.client.DoRawHttpRequest(path, "GET", &bytes.Buffer{})
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: could not SoftLayer_Account#getVirtualDiskImages, error message '%s'", err.Error())
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
