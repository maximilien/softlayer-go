package services_test

import (
	"fmt"
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

		accountService        softlayer.SoftLayer_Account_Service
		virtualGuestService   softlayer.SoftLayer_Virtual_Guest_Service
		productPackageService softlayer.SoftLayer_Product_Package_Service
		networkStorageService softlayer.SoftLayer_Network_Storage_Service
	)

	BeforeEach(func() {
		accountService, err = testhelpers.CreateAccountService()
		Expect(err).ToNot(HaveOccurred())

		virtualGuestService, err = testhelpers.CreateVirtualGuestService()
		Expect(err).ToNot(HaveOccurred())

		productPackageService, err = testhelpers.CreateProductPackageService()
		Expect(err).ToNot(HaveOccurred())

		networkStorageService, err = testhelpers.CreateNetworkStorageService()
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

		It("returns an array of iSCSI network storage", func() {
			networkStorageArray, err := accountService.GetIscsiNetworkStorage()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(networkStorageArray)).To(BeNumerically(">=", 0))
		})

		It("returns an array of SoftLayer_Ssh_Keys objects", func() {
			sshKeys, err := accountService.GetSshKeys()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(sshKeys)).To(BeNumerically(">=", 0))
		})
	})

	Context("uses SoftLayer_ProductPackage to list item prices", func() {
		It("returns an array of SoftLayer_Item_Price under a specific package", func() {
			itemPrices, err := productPackageService.GetItemPrices(0)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(itemPrices)).To(BeNumerically(">=", 0))
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

	FContext("SoftLayer_VirtualGuestService#setUserMetadata and SoftLayer_VirtualGuestService#configureMetadataDisk", func() {
		It("creates ssh key, VirtualGuest, waits for it to be RUNNING, set user data, configures disk, verifies user data, and delete VG", func() {
			sshKeyPath := os.Getenv("SOFTLAYER_GO_TEST_SSH_KEY_PATH2")
			Expect(sshKeyPath).ToNot(Equal(""), "SOFTLAYER_GO_TEST_SSH_KEY_PATH2 env variable is not set")

			createdSshKey := testhelpers.CreateTestSshKey(sshKeyPath)
			testhelpers.WaitForCreatedSshKeyToBePresent(createdSshKey.Id)

			virtualGuest := testhelpers.CreateVirtualGuestAndMarkItTest([]datatypes.SoftLayer_Security_Ssh_Key{createdSshKey})

			testhelpers.WaitForVirtualGuestToBeRunning(virtualGuest.Id)
			testhelpers.WaitForVirtualGuestToHaveNoActiveTransactions(virtualGuest.Id)

			startTime := time.Now()
			userMetadata := "softlayer-go test fake metadata"
			transaction := testhelpers.SetUserMetadataAndConfigureDisk(virtualGuest.Id, userMetadata)
			averageTransactionDuration, err := time.ParseDuration(transaction.TransactionStatus.AverageDuration + "m")
			Ω(err).ShouldNot(HaveOccurred())

			testhelpers.WaitForVirtualGuest(virtualGuest.Id, "RUNNING", averageTransactionDuration)
			testhelpers.WaitForVirtualGuestToHaveNoActiveTransactions(virtualGuest.Id)
			fmt.Printf("====> Set Metadata and configured disk on instance: %d in %d time\n", virtualGuest.Id, time.Since(startTime))

			testUserMetadata(userMetadata)

			startTime = time.Now()
			userMetadata = "softlayer-go test MODIFIED fake metadata"
			testhelpers.SetUserMetadataAndConfigureDisk(virtualGuest.Id, userMetadata)

			testhelpers.WaitForVirtualGuest(virtualGuest.Id, "RUNNING", averageTransactionDuration)
			testhelpers.WaitForVirtualGuestToHaveNoActiveTransactions(virtualGuest.Id)
			fmt.Printf("====> Set Metadata and configured disk on instance: %d in %d time\n", virtualGuest.Id, time.Since(startTime))

			testUserMetadata(userMetadata)
		})
	})

	Context("uses SoftLayer_Network_Storage to manage iSCSI volume", func() {
		It("creates an iSCSI volume and then deletes it", func() {
			iscsiStorage, err := networkStorageService.CreateIscsiVolume(20, "138124")

			Expect(err).ToNot(HaveOccurred())
			Expect(iscsiStorage.Id).ToNot(Equal(0))

			networkStorageService.DeleteIscsiVolume(iscsiStorage.Id, true)
			waitForIscsiStorageToBeDeleted(iscsiStorage.Id)
		})
	})

})

func testUserMetadata(userMetadata string) {
	sshKeyFilePath := os.Getenv("SOFTLAYER_GO_TEST_SSH_KEY_PATH2")
	Expect(sshKeyFilePath).ToNot(Equal(""), "SOFTLAYER_GO_TEST_SSH_KEY_PATH2 env variable is not set")

	workingDir, err := os.Getwd()
	Expect(err).ToNot(HaveOccurred())

	fetchUserMetadataShFilePath := filepath.Join(workingDir, "..", "scripts", "fetch_user_metadata.sh")
	Expect(err).ToNot(HaveOccurred())

	testhelpers.ScpToVirtualGuest(6396994, sshKeyFilePath, fetchUserMetadataShFilePath, "/tmp")
	retCode := testhelpers.SshExecOnVirtualGuest(6396994, sshKeyFilePath, "/tmp/fetch_user_metadata.sh", userMetadata)
	Expect(retCode).To(Equal(0))
}

func waitForIscsiStorageToBeDeleted(storageId int) {
	accountService, err := testhelpers.CreateAccountService()
	Expect(err).ToNot(HaveOccurred())

	fmt.Printf("----> waiting for created iSCSI volume to be deleted\n")
	Eventually(func() bool {
		storages, err := accountService.GetIscsiNetworkStorage()
		Expect(err).ToNot(HaveOccurred())

		deletedFlag := false
		for _, storage := range storages {
			if storage.Id == storageId && storage.BillingItem == nil {
				deletedFlag = true
			}
		}
		return deletedFlag
	}, TIMEOUT, POLLING_INTERVAL).Should(BeTrue(), "created iSCSI volume but not deleted successfully")
}
