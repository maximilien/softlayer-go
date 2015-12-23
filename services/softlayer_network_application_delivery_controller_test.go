package services_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclientfakes "github.com/maximilien/softlayer-go/client/fakes"
	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
	testhelpers "github.com/maximilien/softlayer-go/test_helpers"
)

var _ = Describe("SoftLayer_Network_Application_Delivery_Controller_Service", func() {
	var (
		username, apiKey string

		fakeClient *slclientfakes.FakeSoftLayerClient

		nadcService softlayer.SoftLayer_Network_Application_Delivery_Controller_Service
		err                    error
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		// Use mock for product package service (which provides mock pricing items for "cpu", "ram" and "network speed")
		fakeClient.SoftLayerServices["SoftLayer_Product_Package"] = &testhelpers.MockProductPackageService{}

		nadcService, err = fakeClient.GetSoftLayer_Network_Application_Delivery_Controller_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(nadcService).ToNot(BeNil())
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := nadcService.GetName()
			Expect(name).To(Equal("SoftLayer_Network_Application_Delivery_Controller"))
		})
	})

	Context("#CreateNetscalerVPX", func() {
		BeforeEach(func() {
			fakeClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Network_Application_Delivery_Controller_Service_CreateNetscalerVPX.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a new Netscaler VPX", func() {
			createOptions := softlayer.NetworkApplicationDeliveryControllerCreateOptions {
				Speed: 10,
				Version: "10.1",
				Plan: "Standard",
				IpCount: 2,
				Location: "DALLAS06",
			}

			result, err := nadcService.CreateNetscalerVPX(createOptions)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.CreateDate).To(Equal("2015-12-14T19:24:23+03:00"))
			Expect(result.Id).To(Equal(15293))
			Expect(result.ModifyDate).To(Equal("2015-12-14T19:24:39+03:00"))
			Expect(result.Name).To(Equal("TWCADC795313-1"))
			Expect(result.TypeId).To(Equal(2))
			Expect(result.Type).NotTo(BeNil())
			Expect(result.Type.KeyName).To(Equal("NETSCALER_VPX"))
			Expect(result.Type.Name).To(Equal("NetScaler VPX"))
			Expect(result.Datacenter).NotTo(BeNil())
			Expect(result.Datacenter.Id).To(Equal(154820))
			Expect(result.Datacenter.LongName).To(Equal("Dallas 6"))
			Expect(result.Datacenter.Name).To(Equal("dal06"))
			Expect(result.Description).To(Equal("Citrix NetScaler VPX 10.1 10Mbps Standard"))
			Expect(result.LicenseExpirationDate).To(Equal("2016-09-30T08:00:00+03:00"))
			Expect(result.ManagedResourceFlag).To(Equal(false))
			Expect(result.ManagementIpAddress).To(Equal("10.107.140.243"))
			Expect(result.NetworkVlanCount).To(Equal(1))
			Expect(result.PrimaryIpAddress).To(Equal("184.172.114.147"))
			Expect(result.Password).NotTo(BeNil())
			Expect(result.Password[0].Password).To(Equal("GYdN95kA"))
			Expect(result.Password[0].Username).To(Equal("root"))
		})
	})

	Context("#CreateVirtualIpAddress", func() {
		BeforeEach(func() {
			fakeClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Network_Application_Delivery_Controller_Service_CreateVirtualIpAddress.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a new Virtual Ip Address", func() {
			template := datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress_Template {
				ConnectionLimit: 1,
			}
			nadcId := 123

			result, err := nadcService.CreateVirtualIpAddress(nadcId, template)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(true))
		})
	})

	Context("#GetObject", func() {
		BeforeEach(func() {
			fakeClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Network_Application_Delivery_Controller_GetObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves SoftLayer_Network_Application_Delivery_Controller instance", func() {
			result, err := nadcService.GetObject(111)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Type).To(Equal("someTestType"))
		})
	})

	Context("#GetVirtualIpAddress", func() {
		BeforeEach(func() {
			fakeClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Network_Application_Delivery_Controller_GetVirtualIpAddress.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves SoftLayer_Network_LoadBalancer_VirtualIpAddress instance", func() {
			nadcId := 1234
			vipName := "testVipName"
			result, err := nadcService.GetVirtualIpAddress(nadcId, vipName)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Type).To(Equal("someTestType"))
		})
	})

	Context("#EditVirtualIpAddress", func() {
		BeforeEach(func() {
			fakeClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Dns_Domain_Record_Service_editObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("applies changes to the existing SoftLayer_Dns_Domain_Record instance", func() {
			nadcId := 123
			template := datatypes.SoftLayer_Network_LoadBalancer_VirtualIpAddress_Template {

			}

			result, err := nadcService.EditVirtualIpAddress(nadcId, template)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Data).To(Equal("changedData"))
			Expect(result.DomainId).To(Equal(124))
			Expect(result.Expire).To(Equal(99998))
			Expect(result.Host).To(Equal("changedHost.com"))
			Expect(result.Id).To(Equal(112))
			Expect(result.Minimum).To(Equal(2))
			Expect(result.MxPriority).To(Equal(8))
			Expect(result.Refresh).To(Equal(101))
			Expect(result.ResponsiblePerson).To(Equal("changedTestPerson"))
			Expect(result.Retry).To(Equal(445))
			Expect(result.Ttl).To(Equal(223))
			Expect(result.Type).To(Equal("changedTestType"))
		})
	})

	Context("#DeleteVirtualIpAddress", func() {
		It("sucessfully removes SoftLayer_Network_LoadBalancer_VirtualIpAddress instance from NADC", func() {
			nadcId := 1234
			vipName := "testVipName"
			result, err := nadcService.DeleteVirtualIpAddress(nadcId, vipName)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(true))
		})
	})

	Context("#DeleteObject", func() {
		It("sucessfully removes SoftLayer_Network_Application_Delivery_Controller instance", func() {
			nadcId := 1234
			result, err := nadcService.DeleteObject(nadcId)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(true))
		})
	})

	Context("#getApplicationDeliveryControllerItems", func() {
		BeforeEach(func() {
			fakeClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Product_Order_placeOrder.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("reports error when pricing item for provided CPUs is not available", func() {
			_, err := virtualGuestService.GetAvailableUpgradeItemPrices(&softlayer.UpgradeOptions{Cpus: 3})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Failed to find price for 'cpus' (of size 3)"))
		})

		It("reports error when pricing item for provided RAM is not available", func() {
			_, err := virtualGuestService.GetAvailableUpgradeItemPrices(&softlayer.UpgradeOptions{MemoryInGB: 1500})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Failed to find price for 'memory' (of size 1500)"))
		})

		It("reports error when pricing item for provided network speed is not available", func() {
			_, err := virtualGuestService.GetAvailableUpgradeItemPrices(&softlayer.UpgradeOptions{NicSpeed: 999})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Failed to find price for 'nic_speed' (of size 999)"))
		})
	})
})
