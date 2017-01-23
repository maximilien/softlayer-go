package services_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclientfakes "github.com/maximilien/softlayer-go/client/fakes"
	"github.com/maximilien/softlayer-go/softlayer"
	testhelpers "github.com/maximilien/softlayer-go/test_helpers"
)

var _ = Describe("SoftLayer_Network_Vlan", func() {
	var (
		username, apiKey string

		fakeClient *slclientfakes.FakeSoftLayerClient

		networkVlanService softlayer.SoftLayer_Network_Vlan_Service
		err                error
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		networkVlanService, err = fakeClient.GetSoftLayer_Network_Vlan_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(networkVlanService).ToNot(BeNil())
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := networkVlanService.GetName()
			Expect(name).To(Equal("SoftLayer_Network_Vlan"))
		})
	})

	Context("#GetObject", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Network_Vlan_Service_getObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves SoftLayer_Network_Vlan object", func() {
			networkVlan, err := networkVlanService.GetObject(123456)
			Expect(err).ToNot(HaveOccurred())
			Expect(networkVlan.Id).To(Equal(1234567))
			Expect(networkVlan.Name).To(Equal("Env4 LON02 Fab Pub"))
			Expect(networkVlan.NetworkSpace).To(Equal("PUBLIC"))
			Expect(networkVlan.PrimarySubnetId).To(Equal(123456))
			Expect(networkVlan.VlanNumber).To(Equal(887))
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := networkVlanService.GetObject(123456)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := networkVlanService.GetObject(123456)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})
})
