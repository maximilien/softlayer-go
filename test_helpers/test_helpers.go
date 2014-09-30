package test_helpers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	. "github.com/onsi/gomega"

	slclient "github.com/maximilien/softlayer-go/client"
	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

var (
	TIMEOUT          time.Duration
	POLLING_INTERVAL time.Duration
)

const (
	TEST_NOTES_PREFIX = "TEST:softlayer-go"
	TEST_LABEL_PREFIX = "TEST:softlayer-go"

	MAX_WAIT_RETRIES = 10
	WAIT_TIME        = 5
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
		if key.Notes == TEST_NOTES_PREFIX {
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

func FindAndDeleteTestVirtualGuests() ([]int, error) {
	virtualGuests, err := FindTestVirtualGuests()
	if err != nil {
		return []int{}, err
	}

	virtualGuestService, err := CreateVirtualGuestService()
	if err != nil {
		return []int{}, err
	}

	virtualGuestIds := []int{}
	for _, virtualGuest := range virtualGuests {
		virtualGuestIds = append(virtualGuestIds, virtualGuest.Id)

		deleted, err := virtualGuestService.DeleteObject(virtualGuest.Id)
		if err != nil {
			return []int{}, err
		}

		if !deleted {
			return []int{}, errors.New(fmt.Sprintf("Could not delete virtual guest with id: %d", virtualGuest.Id))
		}
	}

	return virtualGuestIds, nil
}

func MarkVirtualGuestAsTest(virtualGuest datatypes.SoftLayer_Virtual_Guest) error {
	virtualGuestService, err := CreateVirtualGuestService()
	if err != nil {
		return err
	}

	vgTemplate := datatypes.SoftLayer_Virtual_Guest{
		Notes: TEST_NOTES_PREFIX,
	}

	edited, err := virtualGuestService.EditObject(virtualGuest.Id, vgTemplate)
	if err != nil {
		return err
	}
	if edited == false {
		return errors.New(fmt.Sprintf("Could not edit virtual guest with id: %d", virtualGuest.Id))
	}

	return nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return !os.IsNotExist(err)
}

func CreateTestSshKey(sshKeyPath string) datatypes.SoftLayer_Security_Ssh_Key {
	testSshKeyValue, err := ioutil.ReadFile(sshKeyPath)
	Expect(err).ToNot(HaveOccurred())

	sshKey := datatypes.SoftLayer_Security_Ssh_Key{
		Key:         strings.Trim(string(testSshKeyValue), "\n"),
		Fingerprint: "f6:c2:9d:57:2f:74:be:a1:db:71:f2:e5:8e:0f:84:7e",
		Label:       TEST_LABEL_PREFIX,
		Notes:       TEST_NOTES_PREFIX,
	}

	sshKeyService, err := CreateSecuritySshKeyService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> creating ssh key\n")
	createdSshKey, err := sshKeyService.CreateObject(sshKey)
	Expect(err).ToNot(HaveOccurred())
	Expect(createdSshKey.Key).To(Equal(sshKey.Key), "key")
	Expect(createdSshKey.Label).To(Equal(sshKey.Label), "label")
	Expect(createdSshKey.Notes).To(Equal(sshKey.Notes), "notes")
	Expect(createdSshKey.CreateDate).ToNot(BeNil(), "createDate")
	Expect(createdSshKey.Fingerprint).ToNot(Equal(""), "fingerprint")
	Expect(createdSshKey.Id).To(BeNumerically(">", 0), "id")
	Expect(createdSshKey.ModifyDate).To(BeNil(), "modifyDate")
	fmt.Printf("----> created ssh key: %d\n", createdSshKey.Id)

	return createdSshKey
}

func CreateVirtualGuestAndMarkItTest(securitySshKeys []datatypes.SoftLayer_Security_Ssh_Key) datatypes.SoftLayer_Virtual_Guest {
	sshKeys := make([]datatypes.SshKey, len(securitySshKeys))
	for i, securitySshKey := range securitySshKeys {
		sshKeys[i] = datatypes.SshKey{Id: securitySshKey.Id}
	}

	virtualGuestTemplate := datatypes.SoftLayer_Virtual_Guest_Template{
		Hostname:  "test",
		Domain:    "softlayergo.com",
		StartCpus: 1,
		MaxMemory: 1024,
		Datacenter: datatypes.Datacenter{
			Name: "ams01",
		},
		SshKeys:                      sshKeys,
		HourlyBillingFlag:            true,
		LocalDiskFlag:                true,
		OperatingSystemReferenceCode: "UBUNTU_LATEST",
	}

	virtualGuestService, err := CreateVirtualGuestService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> creating new virtual guest\n")
	virtualGuest, err := virtualGuestService.CreateObject(virtualGuestTemplate)
	Expect(err).ToNot(HaveOccurred())
	fmt.Printf("----> created virtual guest: %d\n", virtualGuest.Id)

	WaitForVirtualGuestToBeRunning(virtualGuest.Id)
	WaitForVirtualGuestToHaveNoActiveTransactions(virtualGuest.Id)

	fmt.Printf("----> marking virtual guest with TEST:softlayer-go\n")
	err = MarkVirtualGuestAsTest(virtualGuest)
	Expect(err).ToNot(HaveOccurred(), "Could not mark virtual guest as test")
	fmt.Printf("----> marked virtual guest with TEST:softlayer-go\n")

	return virtualGuest
}

func DeleteVirtualGuest(virtualGuestId int) {
	virtualGuestService, err := CreateVirtualGuestService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> deleting virtual guest: %d\n", virtualGuestId)
	deleted, err := virtualGuestService.DeleteObject(virtualGuestId)
	Expect(err).ToNot(HaveOccurred())
	Expect(deleted).To(BeTrue(), "could not delete virtual guest")

	WaitForVirtualGuestToHaveNoActiveTransactions(virtualGuestId)
}

func DeleteSshKey(sshKeyId int) {
	sshKeyService, err := CreateSecuritySshKeyService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> deleting ssh key: %d\n", sshKeyId)
	deleted, err := sshKeyService.DeleteObject(sshKeyId)
	Expect(err).ToNot(HaveOccurred())
	Expect(deleted).To(BeTrue(), "could not delete ssh key")

	WaitForDeletedSshKeyToNoLongerBePresent(sshKeyId)
}

func WaitForVirtualGuestToBeRunning(virtualGuestId int) {
	virtualGuestService, err := CreateVirtualGuestService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> waiting for virtual guest: %d, until RUNNING\n", virtualGuestId)
	Eventually(func() string {
		vgPowerState, err := virtualGuestService.GetPowerState(virtualGuestId)
		Expect(err).ToNot(HaveOccurred())
		fmt.Printf("----> virtual guest: %d, has power state: %s\n", virtualGuestId, vgPowerState.KeyName)
		return vgPowerState.KeyName
	}, TIMEOUT, POLLING_INTERVAL).Should(Equal("RUNNING"), "failed waiting for virtual guest to be RUNNING")
}

func WaitForVirtualGuestToHaveNoActiveTransactions(virtualGuestId int) {
	virtualGuestService, err := CreateVirtualGuestService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> waiting for virtual guest to have no active transactions pending\n")
	Eventually(func() int {
		activeTransactions, err := virtualGuestService.GetActiveTransactions(virtualGuestId)
		Expect(err).ToNot(HaveOccurred())
		fmt.Printf("----> virtual guest: %d, has %d active transactions\n", virtualGuestId, len(activeTransactions))
		return len(activeTransactions)
	}, TIMEOUT, POLLING_INTERVAL).Should(Equal(0), "failed waiting for virtual guest to have no active transactions")
}

func WaitForDeletedSshKeyToNoLongerBePresent(sshKeyId int) {
	accountService, err := CreateAccountService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> waiting for deleted ssh key to no longer be present\n")
	Eventually(func() bool {
		sshKeys, err := accountService.GetSshKeys()
		Expect(err).ToNot(HaveOccurred())

		deleted := true
		for _, sshKey := range sshKeys {
			if sshKey.Id == sshKeyId {
				deleted = false
			}
		}
		return deleted
	}, TIMEOUT, POLLING_INTERVAL).Should(BeTrue(), "failed waiting for deleted ssh key to be removed from list of ssh keys")
}

func WaitForCreatedSshKeyToBePresent(sshKeyId int) {
	accountService, err := CreateAccountService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> waiting for created ssh key to be present\n")
	Eventually(func() bool {
		sshKeys, err := accountService.GetSshKeys()
		Expect(err).ToNot(HaveOccurred())

		keyPresent := false
		for _, sshKey := range sshKeys {
			if sshKey.Id == sshKeyId {
				keyPresent = true
			}
		}
		return keyPresent
	}, TIMEOUT, POLLING_INTERVAL).Should(BeTrue(), "created ssh key but not in the list of ssh keys")
}
