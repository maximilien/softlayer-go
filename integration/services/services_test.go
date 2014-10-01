package services_test

import (
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
	testhelpers "github.com/maximilien/softlayer-go/test_helpers"
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

		testhelpers.TIMEOUT = 35 * time.Minute
		testhelpers.POLLING_INTERVAL = 10 * time.Second
	})

	Context("SoftLayer_VirtualGuest#<getVirtualDiskImages, getSshKeys, getVirtualGuests, getNetworkStorage>", func() {
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

	Context("SoftLayer_SecuritySshKey#CreateObject and SoftLayer_SecuritySshKey#DeleteObject", func() {
		It("creates the ssh key and verify it is present and then deletes it", func() {
			sshKeyPath := os.Getenv("SOFTLAYER_GO_TEST_SSH_KEY_PATH1")
			Expect(sshKeyPath).ToNot(Equal(""), "SOFTLAYER_GO_TEST_SSH_KEY_PATH1 env variable is not set")

			createdSshKey := testhelpers.CreateTestSshKey(sshKeyPath)
			testhelpers.WaitForCreatedSshKeyToBePresent(createdSshKey.Id)

			sshKeyService, err := testhelpers.CreateSecuritySshKeyService()
			Expect(err).ToNot(HaveOccurred())

			deleted, err := sshKeyService.DeleteObject(createdSshKey.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(deleted).To(BeTrue())

			testhelpers.WaitForDeletedSshKeyToNoLongerBePresent(createdSshKey.Id)
		})
	})

	Context("SoftLayer_VirtualGuest#CreateObject and SoftLayer_VirtualGuest#DeleteObject", func() {
		It("creates the virtual guest instance and waits for it to be active then delete it", func() {
			virtualGuest := testhelpers.CreateVirtualGuestAndMarkItTest([]datatypes.SoftLayer_Security_Ssh_Key{})

			testhelpers.WaitForVirtualGuestToBeRunning(virtualGuest.Id)
			testhelpers.WaitForVirtualGuestToHaveNoActiveTransactions(virtualGuest.Id)

			testhelpers.DeleteVirtualGuest(virtualGuest.Id)
		})
	})

	Context("SoftLayer_SecuritySshKey#CreateObject and SoftLayer_VirtualGuest#CreateObject", func() {
		It("creates key, creates virtual guest and adds key to list of VG", func() {
			sshKeyPath := os.Getenv("SOFTLAYER_GO_TEST_SSH_KEY_PATH2")
			Expect(sshKeyPath).ToNot(Equal(""), "SOFTLAYER_GO_TEST_SSH_KEY_PATH2 env variable is not set")

			createdSshKey := testhelpers.CreateTestSshKey(sshKeyPath)
			testhelpers.WaitForCreatedSshKeyToBePresent(createdSshKey.Id)

			virtualGuest := testhelpers.CreateVirtualGuestAndMarkItTest([]datatypes.SoftLayer_Security_Ssh_Key{createdSshKey})

			testhelpers.WaitForVirtualGuestToBeRunning(virtualGuest.Id)
			testhelpers.WaitForVirtualGuestToHaveNoActiveTransactions(virtualGuest.Id)

			testhelpers.DeleteVirtualGuest(virtualGuest.Id)
			testhelpers.DeleteSshKey(createdSshKey.Id)
		})
	})

	Context("SoftLayer_VirtualGuestService#setUserMetadata and SoftLayer_VirtualGuestService#configureMetadataDisk", func() {
		It("creates ssh key, VirtualGuest, waits for it to be RUNNING, set user data, configures disk, verifies user data, and delete VG", func() {
			sshKeyPath := os.Getenv("SOFTLAYER_GO_TEST_SSH_KEY_PATH2")
			Expect(sshKeyPath).ToNot(Equal(""), "SOFTLAYER_GO_TEST_SSH_KEY_PATH2 env variable is not set")

			createdSshKey := testhelpers.CreateTestSshKey(sshKeyPath)
			testhelpers.WaitForCreatedSshKeyToBePresent(createdSshKey.Id)

			virtualGuest := testhelpers.CreateVirtualGuestAndMarkItTest([]datatypes.SoftLayer_Security_Ssh_Key{createdSshKey})

			testhelpers.WaitForVirtualGuestToBeRunning(virtualGuest.Id)
			testhelpers.WaitForVirtualGuestToHaveNoActiveTransactions(virtualGuest.Id)

			testhelpers.SetUserDataToVirtualGuest(virtualGuest.Id, "softlayer-go test fake metadata")
			transaction := testhelpers.ConfigureMetadataDiskOnVirtualGuest(virtualGuest.Id)

			averageTransactionDuration, err := time.ParseDuration(transaction.TransactionStatus.AverageDuration + "m")
			Expect(err).ToNot(HaveOccurred())

			testhelpers.WaitForVirtualGuest(virtualGuest.Id, "RUNNING", averageTransactionDuration)
			testhelpers.WaitForVirtualGuestToHaveNoActiveTransactions(virtualGuest.Id)

			sshKeyFilePath := os.Getenv("SOFTLAYER_GO_TEST_SSH_KEY_PATH2")
			Expect(sshKeyFilePath).ToNot(Equal(""), "SOFTLAYER_GO_TEST_SSH_KEY_PATH2 env variable is not set")

			workingDir, err := os.Getwd()
			Expect(err).ToNot(HaveOccurred())

			fetchUserMetadataShFilePath := filepath.Join(workingDir, "..", "scripts", "fetch_user_metadata.sh")
			Expect(err).ToNot(HaveOccurred())

			testhelpers.ScpToVirtualGuest(6396994, sshKeyFilePath, fetchUserMetadataShFilePath, "/tmp")
			retCode := testhelpers.SshExecOnVirtualGuest(6396994, sshKeyFilePath, "/tmp/fetch_user_metadata.sh", "softlayer-gp test fake metadata")
			Expect(retCode).To(Equal(0))
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
