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

var _ = Describe("SoftLayer_Ssh_Key_Service", func() {
	var (
		username, apiKey string
		err              error

		fakeClient *slclientfakes.FakeSoftLayerClient

		sshKeyService softlayer.SoftLayer_Security_Ssh_Key_Service

		sshKey         datatypes.SoftLayer_Security_Ssh_Key
		sshKeyTemplate datatypes.SoftLayer_Security_Ssh_Key
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		sshKeyService, err = fakeClient.GetSoftLayer_Security_Ssh_Key_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(sshKeyService).ToNot(BeNil())

		sshKey = datatypes.SoftLayer_Security_Ssh_Key{}
		sshKeyTemplate = datatypes.SoftLayer_Security_Ssh_Key{}
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := sshKeyService.GetName()
			Expect(name).To(Equal("SoftLayer_Security_Ssh_Key"))
		})
	})

	Context("#CreateObject", func() {
		BeforeEach(func() {
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Security_Ssh_Key_Service_createObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a new SoftLayer_Ssh_Key instance", func() {
			sshKeyTemplate = datatypes.SoftLayer_Security_Ssh_Key{
				Fingerprint: "fake-fingerprint",
				Key:         "fake-key",
				Label:       "fake-label",
				Notes:       "fake-notes",
			}
			sshKey, err = sshKeyService.CreateObject(sshKeyTemplate)
			Expect(err).ToNot(HaveOccurred())
			Expect(sshKey.Fingerprint).To(Equal("fake-fingerprint"))
			Expect(sshKey.Key).To(Equal("fake-key"))
			Expect(sshKey.Label).To(Equal("fake-label"))
			Expect(sshKey.Notes).To(Equal("fake-notes"))
		})
	})

	Context("#DeleteObject", func() {
		BeforeEach(func() {
			sshKey.Id = 1234567
		})

		It("sucessfully deletes the SoftLayer_Ssh_Key instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("true")
			deleted, err := sshKeyService.DeleteObject(sshKey.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(deleted).To(BeTrue())
		})

		It("fails to delete the SoftLayer_Ssh_Key instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("false")
			deleted, err := sshKeyService.DeleteObject(sshKey.Id)
			Expect(err).To(HaveOccurred())
			Expect(deleted).To(BeFalse())
		})
	})

	Context("#GetSoftwarePasswords", func() {
		BeforeEach(func() {
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Security_Ssh_Key_Service_getSoftwarePasswords.json")
			Expect(err).ToNot(HaveOccurred())

			sshKey.Id = 1234567
		})

		It("retrieves the software passwords associated with this virtual guest", func() {
			passwords, err := sshKeyService.GetSoftwarePasswords(sshKey.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(passwords)).To(Equal(1))

			password := passwords[0]
			Expect(password.CreateDate).ToNot(BeNil())
			Expect(password.Id).To(Equal(4244148))
			Expect(password.ModifyDate).ToNot(BeNil())
			Expect(password.Notes).To(Equal(""))
			Expect(password.Password).To(Equal("QJ95Gznz"))
			Expect(password.Port).To(Equal(0))
			Expect(password.SoftwareId).To(Equal(4181746))
			Expect(password.Username).To(Equal("root"))

			Expect(password.Software.HardwareId).To(Equal(0))
			Expect(password.Software.Id).To(Equal(4181746))
			Expect(password.Software.ManufacturerLicenseInstance).To(Equal(""))
		})
	})
})
