package services_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	clientfakes "github.com/TheWeatherCompany/softlayer-go/client/fakes"
	"github.com/TheWeatherCompany/softlayer-go/softlayer"
	"github.com/TheWeatherCompany/softlayer-go/test_helpers"
)

var _ = Describe("SoftLayer_User_Customer_Service", func() {
	var (
		username, apiKey    string
		fakeClient          *clientfakes.FakeSoftLayerClient
		userCustomerService softlayer.SoftLayer_User_Customer_Service
		err                 error
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = clientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		userCustomerService, err = fakeClient.GetSoftLayer_User_Customer_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(userCustomerService).ToNot(BeNil())
	})

	Context("#GetApiAuthenticationKeys", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err =
				test_helpers.ReadJsonTestFixtures("services", "SoftLayer_User_Customer_getApiAuthenticationKeys.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an array of datatypes.SoftLayer_User_Customer_ApiAuthentication", func() {
			authKeys, err := userCustomerService.GetApiAuthenticationKeys(0)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(authKeys)).ToNot(Equal(0))
			Expect(authKeys[0].AuthenticationKey).To(Equal("0123456789"))
		})
	})

	Context("#AddApiAuthenticationKey", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err =
				test_helpers.ReadJsonTestFixtures("services", "SoftLayer_User_Customer_addApiAuthenticationKey.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("adds an api authentication key to a softlayer user", func() {
			err := userCustomerService.AddApiAuthenticationKey(0)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
