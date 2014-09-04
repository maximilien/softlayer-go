package test_helpers

import (
	"errors"
	"fmt"
	"os"
	"strings"

	slclient "github.com/maximilien/softlayer-go/client"
	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

const (
	TEST_NOTES_PREFIX = "TEST:softlayer-go"
	TEST_LABEL_PREFIX = "TEST:softlayer-go"
)

func FindTestVirtualGuests() ([]datatypes.SoftLayer_Virtual_Guest, error) {
	accountService, err := CreateAccountService()
	if err != nil {
		return []datatypes.SoftLayer_Virtual_Guest{}, err
	}

	virtualGuests, err := accountService.GetVirtualGuests()
	if err != nil {
		return []datatypes.SoftLayer_Virtual_Guest{}, err
	}

	testVirtualGuests := []datatypes.SoftLayer_Virtual_Guest{}
	for _, vGuest := range virtualGuests {
		if strings.Contains(vGuest.Notes, TEST_NOTES_PREFIX) {
			testVirtualGuests = append(testVirtualGuests, vGuest)
		}
	}

	return testVirtualGuests, nil
}

func FindTestVirtualDiskImages() ([]datatypes.SoftLayer_Virtual_Disk_Image, error) {
	accountService, err := CreateAccountService()
	if err != nil {
		return []datatypes.SoftLayer_Virtual_Disk_Image{}, err
	}

	virtualDiskImages, err := accountService.GetVirtualDiskImages()
	if err != nil {
		return []datatypes.SoftLayer_Virtual_Disk_Image{}, err
	}

	testVirtualDiskImages := []datatypes.SoftLayer_Virtual_Disk_Image{}
	for _, vDI := range virtualDiskImages {
		if strings.Contains(vDI.Description, TEST_NOTES_PREFIX) {
			testVirtualDiskImages = append(testVirtualDiskImages, vDI)
		}
	}

	return testVirtualDiskImages, nil
}

func FindTestNetworkStorage() ([]datatypes.SoftLayer_Network_Storage, error) {
	accountService, err := CreateAccountService()
	if err != nil {
		return []datatypes.SoftLayer_Network_Storage{}, err
	}

	networkStorageArray, err := accountService.GetNetworkStorage()
	if err != nil {
		return []datatypes.SoftLayer_Network_Storage{}, err
	}

	testNetworkStorageArray := []datatypes.SoftLayer_Network_Storage{}
	for _, storage := range networkStorageArray {
		if strings.Contains(storage.Notes, TEST_NOTES_PREFIX) {
			testNetworkStorageArray = append(testNetworkStorageArray, storage)
		}
	}

	return testNetworkStorageArray, nil
}

func FindTestSshKeys() ([]datatypes.SoftLayer_Security_Ssh_Key, error) {
	accountService, err := CreateAccountService()
	if err != nil {
		return []datatypes.SoftLayer_Security_Ssh_Key{}, err
	}

	sshKeys, err := accountService.GetSshKeys()
	if err != nil {
		return []datatypes.SoftLayer_Security_Ssh_Key{}, err
	}

	testSshKeys := []datatypes.SoftLayer_Security_Ssh_Key{}
	for _, key := range sshKeys {
		if strings.Contains(key.Notes, TEST_NOTES_PREFIX) {
			testSshKeys = append(testSshKeys, key)
		}
	}

	return testSshKeys, nil
}

func GetUsernameAndApiKey() (string, string, error) {
	username := os.Getenv("SL_USERNAME")
	if username == "" {
		return "", "", errors.New("SL_USERNAME environment must be set")
	}

	apiKey := os.Getenv("SL_API_KEY")
	if apiKey == "" {
		return username, "", errors.New("SL_API_KEY environment must be set")
	}

	return username, apiKey, nil
}

func CreateAccountService() (softlayer.SoftLayer_Account_Service, error) {
	username, apiKey, err := GetUsernameAndApiKey()
	if err != nil {
		return nil, err
	}

	client := slclient.NewSoftLayerClient(username, apiKey)
	accountService, err := client.GetSoftLayer_Account_Service()
	if err != nil {
		return nil, err
	}

	return accountService, nil
}

func CreateVirtualGuestService() (softlayer.SoftLayer_Virtual_Guest_Service, error) {
	username, apiKey, err := GetUsernameAndApiKey()
	if err != nil {
		return nil, err
	}

	client := slclient.NewSoftLayerClient(username, apiKey)
	virtualGuestService, err := client.GetSoftLayer_Virtual_Guest_Service()
	if err != nil {
		return nil, err
	}

	return virtualGuestService, nil
}

func CreateSecuritySshKeyService() (softlayer.SoftLayer_Security_Ssh_Key_Service, error) {
	username, apiKey, err := GetUsernameAndApiKey()
	if err != nil {
		return nil, err
	}

	client := slclient.NewSoftLayerClient(username, apiKey)
	sshKeyService, err := client.GetSoftLayer_Security_Ssh_Key_Service()
	if err != nil {
		return nil, err
	}

	return sshKeyService, nil
}

func FindAndDeleteTestSshKeys() error {
	sshKeys, err := FindTestSshKeys()
	if err != nil {
		return err
	}

	sshKeyService, err := CreateSecuritySshKeyService()
	if err != nil {
		return err
	}

	for _, sshKey := range sshKeys {
		deleted, err := sshKeyService.DeleteObject(sshKey.Id)
		if err != nil {
			return err
		}
		if !deleted {
			return errors.New(fmt.Sprintf("Could not delete ssh key with id: %d", sshKey.Id))
		}
	}

	return nil
}
