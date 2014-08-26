package services_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclient "github.com/maximilien/softlayer-go/client"
	data_types "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

var _ = Describe("SoftLayer_Virtual_Guest", func() {
	var (
		username, apiKey     string
		client               softlayer.Client
		virtualGuest         softlayer.SoftLayer_Virtual_Guest
		err                  error
		virtualGuestTemplate data_types.SoftLayer_Virtual_Guest_Template
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		client = slclient.NewSoftLayerClient(username, apiKey)
		Expect(client).ToNot(BeNil())

		virtualGuest, err = client.GetSoftLayer_Virtual_Guest()
		Expect(err).ToNot(HaveOccurred())
		Expect(virtualGuest).ToNot(BeNil())
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := virtualGuest.GetName()
			Expect(name).To(Equal("SoftLayer_Virtual_Guest"))
		})
	})

	Context("#CreateObject", func() {
		XIt("creates a new SoftLayer_Virtual_Guest instance", func() {
			virtualGuest, err := virtualGuest.CreateObject(virtualGuestTemplate)
			Expect(err).ToNot(HaveOccurred())
			Expect(virtualGuest).ToNot(BeNil())
		})

		// It("generates the correct JSON body for the SoftLayer_Virtual_Guest/createObject.json POST call", func() {

		// 	})

		It("flags all missing required parameters for SoftLayer_Virtual_Guest/createObject.json POST call", func() {

		})

	})

	Context("#DeleteObject", func() {
		XIt("deletes the SoftLayer_Virtual_Guest", func() {
			deleted, err := virtualGuest.DeleteObject(virtualGuestTemplate)
			Expect(err).ToNot(HaveOccurred())
			Expect(deleted).To(BeTrue())
		})
	})
})
