package services_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclientfakes "github.com/maximilien/softlayer-go/client/fakes"
	common "github.com/maximilien/softlayer-go/common"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

var _ = Describe("SoftLayer_Network_Storage", func() {
	var (
		username, apiKey string

		fakeClient *slclientfakes.FakeSoftLayerClient

		networkStorageService softlayer.SoftLayer_Network_Storage_Service
		err                   error
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		networkStorageService, err = fakeClient.GetSoftLayer_Network_Storage_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(networkStorageService).ToNot(BeNil())
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := networkStorageService.GetName()
			Expect(name).To(Equal("SoftLayer_Network_Storage"))
		})
	})

	Context("#GetIscsiVolume", func() {
		BeforeEach(func() {
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Network_Storage_Service_getIscsiVolume.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns the iSCSI volume object based on volume id", func() {
			volume, err := networkStorageService.GetIscsiVolume(1)
			Expect(err).ToNot(HaveOccurred())
			Expect(volume.Id).To(Equal(1))
			Expect(volume.Username).To(Equal("test_username"))
			Expect(volume.Password).To(Equal("test_password"))
			Expect(volume.CapacityGb).To(Equal(20))
		})
	})

})
