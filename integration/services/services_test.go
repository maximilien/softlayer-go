package services_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
	testhelpers "github.com/maximilien/softlayer-go/test_helpers"
)

var (
	TIMEOUT          time.Duration
	POLLING_INTERVAL time.Duration
)

var _ = Describe("SoftLayer Services", func() {
	var (
		err error

		accountService      softlayer.SoftLayer_Account_Service
		virtualGuestService softlayer.SoftLayer_Virtual_Guest_Service
	)

	BeforeEach(func() {
		accountService, err = testhelpers.CreateAccountService()
		Expect(err).ToNot(HaveOccurred())

		virtualGuestService, err = testhelpers.CreateVirtualGuestService()
		Expect(err).ToNot(HaveOccurred())

		TIMEOUT = 25 * time.Minute
		POLLING_INTERVAL = 15 * time.Second
	})

	Context("uses SoftLayer_Account to list current virtual: disk images, guests, ssh keys, and network storage", func() {
		It("returns an array of SoftLayer_Virtual_Guest disk images", func() {
			virtualDiskImages, err := accountService.GetVirtualDiskImages()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(virtualDiskImages)).To(BeNumerically(">=", 0))
		})

		It("returns an array of SoftLayer_Virtual_Guest objects", func() {
			virtualGuests, err := accountService.GetVirtualGuests()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(virtualGuests)).To(BeNumerically(">=", 0))
		})

		It("returns an array of SoftLayer_Virtual_Guest network storage", func() {
			networkStorageArray, err := accountService.GetNetworkStorage()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(networkStorageArray)).To(BeNumerically(">=", 0))
		})

		It("returns an array of SoftLayer_Ssh_Keys objects", func() {
			sshKeys, err := accountService.GetSshKeys()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(sshKeys)).To(BeNumerically(">=", 0))
		})
	})

	Context("uses SoftLayer_Account to create and then delete a an ssh key", func() {
		It("creates the ssh key and verify it is present and then deletes it", func() {
			sshKeyPath := os.Getenv("SOFTLAYER_GO_TEST_SSH_KEY_PATH1")
			Expect(sshKeyPath).ToNot(Equal(""), "SOFTLAYER_GO_TEST_SSH_KEY_PATH1 env variable is not set")

			createdSshKey := createTestSshKey(sshKeyPath)
			waitForCreatedSshKeyToBePresent(createdSshKey.Id)

			sshKeyService, err := testhelpers.CreateSecuritySshKeyService()
			Expect(err).ToNot(HaveOccurred())

			deleted, err := sshKeyService.DeleteObject(createdSshKey.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(deleted).To(BeTrue())

			waitForDeletedSshKeyToNoLongerBePresent(createdSshKey.Id)
		})
	})

	Context("uses SoftLayer_Account to create and then delete a virtual guest instance", func() {
		It("creates the virtual guest instance and waits for it to be active then delete it", func() {
			virtualGuest := createVirtualGuestAndMarkItTest([]datatypes.SoftLayer_Security_Ssh_Key{})

			waitForVirtualGuestToBeRunning(virtualGuest.Id)
			waitForVirtualGuestToHaveNoActiveTransactions(virtualGuest.Id)

			deleteVirtualGuest(virtualGuest.Id)
		})
	})

	Context("uses SoftLayer_Account to create ssh key and new virtual guest with ssh key assigned", func() {
		It("creates key, creates virtual guest and adds key to list of VG", func() {
			sshKeyPath := os.Getenv("SOFTLAYER_GO_TEST_SSH_KEY_PATH2")
			Expect(sshKeyPath).ToNot(Equal(""), "SOFTLAYER_GO_TEST_SSH_KEY_PATH2 env variable is not set")

			createdSshKey := createTestSshKey(sshKeyPath)
			waitForCreatedSshKeyToBePresent(createdSshKey.Id)

			virtualGuest := createVirtualGuestAndMarkItTest([]datatypes.SoftLayer_Security_Ssh_Key{createdSshKey})

			waitForVirtualGuestToBeRunning(virtualGuest.Id)
			waitForVirtualGuestToHaveNoActiveTransactions(virtualGuest.Id)

			deleteVirtualGuest(virtualGuest.Id)
			deleteSshKey(createdSshKey.Id)
		})
	})

	XContext("uses SoftLayer_Account to create a new instance and network storage and attach them", func() {
		It("creates the virtual guest instance and waits for it to be active", func() {
			Expect(false).To(BeTrue())
		})

		It("creates the disk storage and attaches it to the instance", func() {
			Expect(false).To(BeTrue())
		})

		It("deletes the virtual guest instance if it is running", func() {
			Expect(false).To(BeTrue())
		})

		It("detaches and deletes the network storage if available", func() {
			Expect(false).To(BeTrue())
		})
	})
})

func createTestSshKey(sshKeyPath string) datatypes.SoftLayer_Security_Ssh_Key {
	testSshKeyValue, err := ioutil.ReadFile(sshKeyPath)
	Expect(err).ToNot(HaveOccurred())

	sshKey := datatypes.SoftLayer_Security_Ssh_Key{
		Key:         strings.Trim(string(testSshKeyValue), "\n"),
		Fingerprint: "f6:c2:9d:57:2f:74:be:a1:db:71:f2:e5:8e:0f:84:7e",
		Label:       testhelpers.TEST_LABEL_PREFIX,
		Notes:       testhelpers.TEST_NOTES_PREFIX,
	}

	sshKeyService, err := testhelpers.CreateSecuritySshKeyService()
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

func createVirtualGuestAndMarkItTest(securitySshKeys []datatypes.SoftLayer_Security_Ssh_Key) datatypes.SoftLayer_Virtual_Guest {
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

	virtualGuestService, err := testhelpers.CreateVirtualGuestService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> creating new virtual guest\n")
	virtualGuest, err := virtualGuestService.CreateObject(virtualGuestTemplate)
	Expect(err).ToNot(HaveOccurred())
	fmt.Printf("----> created virtual guest: %d\n", virtualGuest.Id)

	waitForVirtualGuestToBeRunning(virtualGuest.Id)
	waitForVirtualGuestToHaveNoActiveTransactions(virtualGuest.Id)

	fmt.Printf("----> marking virtual guest with TEST:softlayer-go\n")
	err = testhelpers.MarkVirtualGuestAsTest(virtualGuest)
	Expect(err).ToNot(HaveOccurred(), "Could not mark virtual guest as test")
	fmt.Printf("----> marked virtual guest with TEST:softlayer-go\n")

	return virtualGuest
}

func deleteVirtualGuest(virtualGuestId int) {
	virtualGuestService, err := testhelpers.CreateVirtualGuestService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> deleting virtual guest: %d\n", virtualGuestId)
	deleted, err := virtualGuestService.DeleteObject(virtualGuestId)
	Expect(err).ToNot(HaveOccurred())
	Expect(deleted).To(BeTrue(), "could not delete virtual guest")

	waitForVirtualGuestToHaveNoActiveTransactions(virtualGuestId)
}

func deleteSshKey(sshKeyId int) {
	sshKeyService, err := testhelpers.CreateSecuritySshKeyService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> deleting ssh key: %d\n", sshKeyId)
	deleted, err := sshKeyService.DeleteObject(sshKeyId)
	Expect(err).ToNot(HaveOccurred())
	Expect(deleted).To(BeTrue(), "could not delete ssh key")

	waitForDeletedSshKeyToNoLongerBePresent(sshKeyId)
}

func waitForVirtualGuestToBeRunning(virtualGuestId int) {
	virtualGuestService, err := testhelpers.CreateVirtualGuestService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> waiting for virtual guest: %d, until RUNNING\n", virtualGuestId)
	Eventually(func() string {
		vgPowerState, err := virtualGuestService.GetPowerState(virtualGuestId)
		Expect(err).ToNot(HaveOccurred())
		fmt.Printf("----> virtual guest: %d, has power state: %s\n", virtualGuestId, vgPowerState.KeyName)
		return vgPowerState.KeyName
	}, TIMEOUT, POLLING_INTERVAL).Should(Equal("RUNNING"), "failed waiting for virtual guest to be RUNNING")
}

func waitForVirtualGuestToHaveNoActiveTransactions(virtualGuestId int) {
	virtualGuestService, err := testhelpers.CreateVirtualGuestService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> waiting for virtual guest to have no active transactions pending\n")
	Eventually(func() int {
		activeTransactions, err := virtualGuestService.GetActiveTransactions(virtualGuestId)
		Expect(err).ToNot(HaveOccurred())
		fmt.Printf("----> virtual guest: %d, has %d active transactions\n", virtualGuestId, len(activeTransactions))
		return len(activeTransactions)
	}, TIMEOUT, POLLING_INTERVAL).Should(Equal(0), "failed waiting for virtual guest to have no active transactions")
}

func waitForDeletedSshKeyToNoLongerBePresent(sshKeyId int) {
	accountService, err := testhelpers.CreateAccountService()
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

func waitForCreatedSshKeyToBePresent(sshKeyId int) {
	accountService, err := testhelpers.CreateAccountService()
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
