package services_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclientfakes "github.com/maximilien/softlayer-go/client/fakes"
	common "github.com/maximilien/softlayer-go/common"
	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

var _ = Describe("SoftLayer_Virtual_Guest_Service", func() {
	var (
		username, apiKey string
		err              error

		fakeClient *slclientfakes.FakeSoftLayerClient

		virtualGuestService softlayer.SoftLayer_Virtual_Guest_Service

		virtualGuest         datatypes.SoftLayer_Virtual_Guest
		virtualGuestTemplate datatypes.SoftLayer_Virtual_Guest_Template
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		virtualGuestService, err = fakeClient.GetSoftLayer_Virtual_Guest_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(virtualGuestService).ToNot(BeNil())

		virtualGuest = datatypes.SoftLayer_Virtual_Guest{}
		virtualGuestTemplate = datatypes.SoftLayer_Virtual_Guest_Template{}
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := virtualGuestService.GetName()
			Expect(name).To(Equal("SoftLayer_Virtual_Guest"))
		})
	})

	Context("#CreateObject", func() {
		BeforeEach(func() {
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_createObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a new SoftLayer_Virtual_Guest instance", func() {
			virtualGuestTemplate = datatypes.SoftLayer_Virtual_Guest_Template{
				Hostname:  "fake-hostname",
				Domain:    "fake.domain.com",
				StartCpus: 2,
				MaxMemory: 1024,
				Datacenter: datatypes.Datacenter{
					Name: "fake-datacenter-name",
				},
				HourlyBillingFlag:            true,
				LocalDiskFlag:                false,
				DedicatedAccountHostOnlyFlag: false,
			}
			virtualGuest, err = virtualGuestService.CreateObject(virtualGuestTemplate)
			Expect(err).ToNot(HaveOccurred())
			Expect(virtualGuest.Hostname).To(Equal("fake-hostname"))
			Expect(virtualGuest.Domain).To(Equal("fake.domain.com"))
			Expect(virtualGuest.StartCpus).To(Equal(2))
			Expect(virtualGuest.MaxMemory).To(Equal(1024))
			Expect(virtualGuest.DedicatedAccountHostOnlyFlag).To(BeFalse())
		})

		It("flags all missing required parameters for SoftLayer_Virtual_Guest/createObject.json POST call", func() {
			virtualGuestTemplate = datatypes.SoftLayer_Virtual_Guest_Template{}
			_, err := virtualGuestService.CreateObject(virtualGuestTemplate)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Hostname"))
			Expect(err.Error()).To(ContainSubstring("Domain"))
			Expect(err.Error()).To(ContainSubstring("StartCpus"))
			Expect(err.Error()).To(ContainSubstring("MaxMemory"))
			Expect(err.Error()).To(ContainSubstring("Datacenter"))
		})
	})

	Context("#DeleteObject", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
		})

		It("sucessfully deletes the SoftLayer_Virtual_Guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("true")
			deleted, err := virtualGuestService.DeleteObject(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(deleted).To(BeTrue())
		})

		It("fails to delete the SoftLayer_Virtual_Guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("false")
			deleted, err := virtualGuestService.DeleteObject(virtualGuest.Id)
			Expect(err).To(HaveOccurred())
			Expect(deleted).To(BeFalse())
		})
	})

	Context("#GetPowerState", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_getPowerState.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves SoftLayer_Virtual_Guest_State for RUNNING instance", func() {
			vgPowerState, err := virtualGuestService.GetPowerState(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(vgPowerState.KeyName).To(Equal("RUNNING"))
		})
	})
})
