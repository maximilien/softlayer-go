package virtual_guest_lifecycle_test

import (
	"fmt"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
	testhelpers "github.com/maximilien/softlayer-go/test_helpers"
)

var _ = Describe("SoftLayer Network Storage Lifecycle", func() {
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

	Context("SoftLayer_Account#<getVirtualDiskImages, getNetworkStorage>", func() {
		It("returns an array of SoftLayer_Virtual_Guest disk images", func() {
			virtualDiskImages, err := accountService.GetVirtualDiskImages()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(virtualDiskImages)).To(BeNumerically(">=", 0))
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

		It("returns an array of SoftLayer_Virtual_Guest_Block_Device_Template_Group objects", func() {
			groups, err := accountService.GetBlockDeviceTemplateGroups()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(groups)).To(BeNumerically(">=", 0))
		})
	})

	Context("uses SoftLayer_ProductPackage to list item prices", func() {
		It("returns an array of SoftLayer_Item_Price under a specific package", func() {
			itemPrices, err := productPackageService.GetItemPrices(0)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(itemPrices)).To(BeNumerically(">=", 0))
		})
	})

	Context("uses SoftLayer_Network_Storage & SoftLayer_Virtual_Guest to manage iSCSI volume", func() {
		It("create a virutal guest and an iSCSI volume, then attaches the iSCSI volume to virtual guest and detaches it, finally delete them.", func() {
			virtualGuest := testhelpers.CreateVirtualGuestAndMarkItTest([]datatypes.SoftLayer_Security_Ssh_Key{})

			testhelpers.WaitForVirtualGuestToBeRunning(virtualGuest.Id)
			testhelpers.WaitForVirtualGuestToHaveNoActiveTransactions(virtualGuest.Id)

			iscsiStorage, err := networkStorageService.CreateIscsiVolume(20, "138124")
			Expect(err).ToNot(HaveOccurred())
			Expect(iscsiStorage.Id).ToNot(Equal(0))

			iscsiVolume, err := networkStorageService.GetIscsiVolume(iscsiStorage.Id)

			Expect(err).ToNot(HaveOccurred())
			Expect(iscsiVolume.Id).To(Equal(iscsiStorage.Id))
			Expect(iscsiVolume.CapacityGb).To(Equal(iscsiStorage.CapacityGb))

			device, err := virtualGuestService.AttachIscsiVolume(virtualGuest.Id, iscsiStorage.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(device).To(Equal("sda"))

			err = virtualGuestService.DetachIscsiVolume(virtualGuest.Id, iscsiStorage.Id)
			Expect(err).ToNot(HaveOccurred())

			networkStorageService.DeleteIscsiVolume(iscsiStorage.Id, true)
			testhelpers.WaitForIscsiStorageToBeDeleted(iscsiStorage.Id)

			testhelpers.DeleteVirtualGuest(virtualGuest.Id)
		})
	})
})
