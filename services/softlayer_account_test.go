package services_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclient "github.com/maximilien/softlayer-go/client"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

var _ = Describe("SoftLayer_Account", func() {
	var (
		username, apiKey string
		client           softlayer.Client
		account          softlayer.SoftLayer_Account
		err              error
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		client = slclient.NewSoftLayerClient(username, apiKey)
		Expect(client).ToNot(BeNil())

		account, err = client.GetSoftLayer_Account()
		Expect(err).To(BeNil())
		Expect(account).ToNot(BeNil())
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := account.GetName()
			Expect(name).To(Equal("SoftLayer_Account"))
		})
	})

	Context("#GetVirtualGuests", func() {
		It("returns an array of datatypes.SoftLayer_Virtual_Guest", func() {
			_, err := account.GetVirtualGuests()
			Expect(err).To(BeNil())
		})
	})

	Context("#GetNetworkStorage", func() {
		It("returns an array of datatypes.SoftLayer_Network_Storage", func() {
			_, err := account.GetNetworkStorage()
			Expect(err).To(BeNil())
		})
	})
})
