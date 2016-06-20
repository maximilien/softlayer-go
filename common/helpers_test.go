package common_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclientfakes "github.com/TheWeatherCompany/softlayer-go/client/fakes"
	"github.com/TheWeatherCompany/softlayer-go/common"
	testhelpers "github.com/TheWeatherCompany/softlayer-go/test_helpers"
)

var _ = Describe("SoftlayerLookupHelpers", func() {
	var (
		username, apiKey string

		fakeClient *slclientfakes.FakeSoftLayerClient
		err        error
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())
	})

	Context("#GetDatacenterByName", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetDatacenterByName.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves ID of datacenter", func() {
			id, err := common.GetDatacenterByName(fakeClient, "ams01")
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal(265592))
		})
	})

	Context("#GetRoutingTypesByName", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetRoutingTypeByName.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves ID of routing type", func() {
			id, err := common.GetRoutingTypeByName(fakeClient, "DNS")
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal(4))
		})
	})

	Context("#GetRoutingMethodsByName", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetRoutingMethodByName.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves ID of routing type", func() {
			id, err := common.GetRoutingMethodByName(fakeClient, "ROUND_ROBIN")
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal(10))
		})
	})

	Context("#GetHealthCheckTypesByName", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetHealthCheckTypeByName.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves ID of health check type", func() {
			id, err := common.GetHealthCheckTypeByName(fakeClient, "DNS")
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal(3))
		})
	})
})
