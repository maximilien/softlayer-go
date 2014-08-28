package client_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclient "github.com/maximilien/softlayer-go/client"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

var _ = Describe("SoftLayerClient", func() {
	var (
		username string
		apiKey   string

		client softlayer.Client
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		apiKey = os.Getenv("SL_API_KEY")

		client = slclient.NewSoftLayerClient(username, apiKey)
	})

	Context("#NewSoftLayerClient", func() {
		It("creates a new client with username and apiKey", func() {
			Expect(username).ToNot(Equal(""), "username cannot be empty, set SL_USERNAME")
			Expect(apiKey).ToNot(Equal(""), "apiKey cannot be empty, set SL_API_KEY")

			client = slclient.NewSoftLayerClient(username, apiKey)
			Expect(client).ToNot(BeNil())
		})
	})

	Context("#GetService", func() {
		It("returns a service with name specified", func() {
			accountService, err := client.GetService("SoftLayer_Account_Service")
			Expect(err).ToNot(HaveOccurred())
			Expect(accountService).ToNot(BeNil())
		})

		It("fails when passed a bad service name", func() {
			_, err := client.GetService("fake-service-name")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("softlayer-go does not support service 'fake-service-name'"))
		})
	})

	Context("#GetSoftLayer_Account", func() {
		It("returns a instance implemementing the SoftLayer_Account interface", func() {
			var accountService softlayer.SoftLayer_Account_Service
			accountService, err := client.GetSoftLayer_Account_Service()
			Expect(err).ToNot(HaveOccurred())
			Expect(accountService).ToNot(BeNil())
		})
	})
})
