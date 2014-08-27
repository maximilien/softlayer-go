package services_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclient "github.com/maximilien/softlayer-go/client"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

var _ = Describe("SoftLayer_Account_Service", func() {
	var (
		username, apiKey string
		client           softlayer.Client
		accountService   softlayer.SoftLayer_Account_Service
		err              error
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		client = slclient.NewSoftLayerClient(username, apiKey)
		Expect(client).ToNot(BeNil())

		accountService, err = client.GetSoftLayer_Account_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(accountService).ToNot(BeNil())
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := accountService.GetName()
			Expect(name).To(Equal("SoftLayer_Account"))
		})
	})

	Context("#GetAccountStatus", func() {
		It("returns an instance of datatypes.SoftLayer_Account_Status that is Active", func() {
			accountStatus, err := accountService.GetAccountStatus()
			Expect(err).ToNot(HaveOccurred())
			Expect(accountStatus.Id).ToNot(Equal(0))
			Expect(accountStatus.Name).To(Equal("Active"))
		})
	})

	Context("#GetVirtualGuests", func() {
		It("returns an array of datatypes.SoftLayer_Virtual_Guest", func() {
			virtualGuests, err := accountService.GetVirtualGuests()
			Expect(err).ToNot(HaveOccurred())
			Expect(virtualGuests).ToNot(BeNil())
		})
	})

	Context("#GetNetworkStorage", func() {
		It("returns an array of datatypes.SoftLayer_Network_Storage", func() {
			networkStorage, err := accountService.GetNetworkStorage()
			Expect(err).ToNot(HaveOccurred())
			Expect(networkStorage).ToNot(BeNil())
		})
	})

	Context("#GetVirtualDiskImages", func() {
		It("returns an array of datatypes.SoftLayer_Virtual_Disk_Image", func() {
			virtualDiskImages, err := accountService.GetVirtualDiskImages()
			Expect(err).ToNot(HaveOccurred())
			Expect(virtualDiskImages).ToNot(BeNil())
		})
	})
})
