package client_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	client "github.com/maximilien/softlayer-go/client"
)

var _ = Describe("softLayerClient", func() {
	var (
		username string
		apiKey   string
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		apiKey = os.Getenv("SL_API_KEY")
	})

	Context("NewSoftLayerClient", func() {
		It("creates a new client with username and apiKey", func() {
			Expect(username).ToNot(Equal(""), "username cannot be empty, set SL_USERNAME")
			Expect(apiKey).ToNot(Equal(""), "apiKey cannot be empty, set SL_API_KEY")

			client := client.NewSoftLayerClient(username, apiKey)
			Expect(client).ToNot(BeNil())
		})
	})

})
