package services_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclientfakes "github.com/TheWeatherCompany/softlayer-go/client/fakes"
	"github.com/TheWeatherCompany/softlayer-go/data_types"
	"github.com/TheWeatherCompany/softlayer-go/softlayer"
	"github.com/TheWeatherCompany/softlayer-go/test_helpers"
)

var _ = Describe("SoftLayer_Billing_Item", func() {
	var (
		username, apiKey string

		fakeClient *slclientfakes.FakeSoftLayerClient

		billingItemService softlayer.SoftLayer_Billing_Item_Service
		err                error
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		billingItemService, err = fakeClient.GetSoftLayer_Billing_Item_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(billingItemService).ToNot(BeNil())
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := billingItemService.GetName()
			Expect(name).To(Equal("SoftLayer_Billing_Item"))
		})
	})

	Context("#CancelService", func() {

		It("returns true", func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse = []byte("true")
			deleted, err := billingItemService.CancelService(1234567)
			Expect(err).ToNot(HaveOccurred())
			Expect(deleted).To(BeTrue())
		})

		It("returns false", func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse = []byte("false")
			deleted, err := billingItemService.CancelService(1234567)
			Expect(err).ToNot(HaveOccurred())
			Expect(deleted).To(BeFalse())
		})
	})

	Context("#CheckOrderStatus", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err =
				test_helpers.ReadJsonTestFixtures("services", "SoftLayer_Billing_Item_Check_Order_Status.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns a boolean when checking an order status", func() {
			completed, _, err := billingItemService.CheckOrderStatus(&data_types.SoftLayer_Container_Product_Order_Receipt{
				PlacedOrder: data_types.SoftLayer_Billing_Order{
					Items: []data_types.SoftLayer_Billing_Order_Item{
						data_types.SoftLayer_Billing_Order_Item{Id: 123456789},
					},
				},
			}, "COMPLETE")
			Expect(err).ToNot(HaveOccurred())
			Expect(completed).To(BeTrue())
		})
	})
})
