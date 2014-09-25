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

	Context("#EditObject", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_editObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("edits an existing SoftLayer_Virtual_Guest instance", func() {
			virtualGuest := datatypes.SoftLayer_Virtual_Guest{
				Notes: "fake-notes",
			}
			edited, err := virtualGuestService.EditObject(virtualGuest.Id, virtualGuest)
			Expect(err).ToNot(HaveOccurred())
			Expect(edited).To(BeTrue())
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

	Context("#GetActiveTransaction", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_getActiveTransaction.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves SoftLayer_Provisioning_Version1_Transaction for virtual guest", func() {
			activeTransaction, err := virtualGuestService.GetActiveTransaction(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(activeTransaction.CreateDate).ToNot(BeNil())
			Expect(activeTransaction.ElapsedSeconds).To(BeNumerically(">", 0))
			Expect(activeTransaction.GuestId).To(Equal(virtualGuest.Id))
			Expect(activeTransaction.Id).To(BeNumerically(">", 0))
		})
	})

	Context("#GetActiveTransactions", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_getActiveTransactions.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves an array of SoftLayer_Provisioning_Version1_Transaction for virtual guest", func() {
			activeTransactions, err := virtualGuestService.GetActiveTransactions(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(activeTransactions)).To(BeNumerically(">", 0))

			for _, activeTransaction := range activeTransactions {
				Expect(activeTransaction.CreateDate).ToNot(BeNil())
				Expect(activeTransaction.ElapsedSeconds).To(BeNumerically(">", 0))
				Expect(activeTransaction.GuestId).To(Equal(virtualGuest.Id))
				Expect(activeTransaction.Id).To(BeNumerically(">", 0))
			}
		})
	})

	Context("#GetSshKeys", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_getSshKeys.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves an array of SoftLayer_Security_Ssh_Key for virtual guest", func() {
			sshKeys, err := virtualGuestService.GetSshKeys(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(sshKeys)).To(BeNumerically(">", 0))

			for _, sshKey := range sshKeys {
				Expect(sshKey.CreateDate).ToNot(BeNil())
				Expect(sshKey.Fingerprint).To(Equal("f6:c2:9d:57:2f:74:be:a1:db:71:f2:e5:8e:0f:84:7e"))
				Expect(sshKey.Id).To(Equal(84386))
				Expect(sshKey.Key).ToNot(Equal(""))
				Expect(sshKey.Label).To(Equal("TEST:softlayer-go"))
				Expect(sshKey.ModifyDate).To(BeNil())
				Expect(sshKey.Label).To(Equal("TEST:softlayer-go"))
			}
		})
	})

	Context("#SetUserMetadata", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_setUserMetadata.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully adds metadata strings as a dile to virtual guest's metadata disk", func() {
			retBool, err := virtualGuestService.SetMetadata(virtualGuest.Id, "fake-metadata")
			Expect(err).ToNot(HaveOccurred())

			Expect(retBool).To(BeTrue())
		})
	})

	Context("#ConfigureMetadataDisk", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_configureMetadataDisk.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully configures a metadata disk for a virtual guest", func() {
			transaction, err := virtualGuestService.ConfigureMetadataDisk(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())

			Expect(transaction.CreateDate).ToNot(BeNil())
			Expect(transaction.ElapsedSeconds).To(Equal(0))
			Expect(transaction.GuestId).To(Equal(virtualGuest.Id))
			Expect(transaction.HardwareId).To(Equal(0))
			Expect(transaction.Id).To(Equal(12476326))
			Expect(transaction.ModifyDate).ToNot(BeNil())
			Expect(transaction.StatusChangeDate).ToNot(BeNil())

			Expect(transaction.TransactionGroup.AverageTimeToComplete).To(Equal("1.62"))
			Expect(transaction.TransactionGroup.Name).To(Equal("Configure Cloud Metadata Disk"))

			Expect(transaction.TransactionStatus.AverageDuration).To(Equal(".32"))
			Expect(transaction.TransactionStatus.FriendlyName).To(Equal("Configure Cloud Metadata Disk"))
			Expect(transaction.TransactionStatus.Name).To(Equal("CLOUD_CONFIGURE_METADATA_DISK"))
		})
	})
})
