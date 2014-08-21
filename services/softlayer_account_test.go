package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclient "github.com/maximilien/softlayer-go/client"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

var _ = Describe("SoftLayer_Account", func() {
	var (
		client  softlayer.Client
		account softlayer.SoftLayer_Account
	)

	BeforeEach(func() {
		client = slclient.NewSoftLayerClient("fake-username", "fake-api-key")
		Expect(client).ToNot(BeNil())

		account, err := client.GetSoftLayer_Account()
		Expect(err).ToNot(BeNil())
		Expect(account).ToNot(BeNil())
	})

	Context("GetName", func() {
		It("returns the name for the service", func() {
			name := account.GetName()
			Expect(name).To(Equal("SoftLayer_Account"))
		})
	})
})
