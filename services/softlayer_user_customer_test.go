package services_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	clientfakes "github.com/TheWeatherCompany/softlayer-go/client/fakes"
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
	"github.com/TheWeatherCompany/softlayer-go/softlayer"
	"github.com/TheWeatherCompany/softlayer-go/test_helpers"
)

var _ = Describe("SoftLayer_User_Customer_Service", func() {
	var (
		username, apiKey    string
		fakeClient          *clientfakes.FakeSoftLayerClient
		userCustomerService softlayer.SoftLayer_User_Customer_Service
		err                 error

		password         string
		permissions      []string
		userCustObj      datatypes.SoftLayer_User_Customer
		userCustInputObj datatypes.SoftLayer_User_Customer
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = clientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		userCustObj = datatypes.SoftLayer_User_Customer{}
		permissions = []string{"ACCESS_ALL_HARDWARE", "ACCESS_ALL_GUEST"}
		userCustInputObj = datatypes.SoftLayer_User_Customer{
			Address1:    "555 Bailey Ave",
			City:        "San Jose",
			CompanyName: "TWC an IBM company",
			Country:     "US",
			Email:       "user4testing@twc.ibm.com",
			FirstName:   "TestUserFirstName",
			LastName:    "TestUserLastName",
			Permissions: []datatypes.SoftLayer_User_Customer_CustomerPermission_Permission{
				datatypes.SoftLayer_User_Customer_CustomerPermission_Permission{
					KeyName: permissions[0],
				},
				datatypes.SoftLayer_User_Customer_CustomerPermission_Permission{
					KeyName: permissions[1],
				},
			},
			State:      "CA",
			Timezone:   107,
			UserStatus: 1001,
			Username:   "user4testing_01",
		}
		password = "Change3Me!" // Needs to conform to SoftLayer password policies.

		userCustomerService, err = fakeClient.GetSoftLayer_User_Customer_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(userCustomerService).ToNot(BeNil())
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := userCustomerService.GetName()
			Expect(name).To(Equal("SoftLayer_User_Customer"))
		})
	})

	Context("#CreateObject", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err =
				test_helpers.ReadJsonTestFixtures("services", "SoftLayer_User_Customer_createAndGetObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a new User_Customer object", func() {
			userCustObj, err = userCustomerService.CreateObject(userCustInputObj, password)
			Expect(err).ToNot(HaveOccurred())
			Expect(userCustObj.Address1).To(Equal(userCustInputObj.Address1))
			Expect(userCustObj.City).To(Equal(userCustInputObj.City))
			Expect(userCustObj.CompanyName).To(Equal(userCustInputObj.CompanyName))
			Expect(userCustObj.Country).To(Equal(userCustInputObj.Country))
			Expect(userCustObj.Email).To(Equal(userCustInputObj.Email))
			Expect(userCustObj.FirstName).To(Equal(userCustInputObj.FirstName))
			Expect(userCustObj.LastName).To(Equal(userCustInputObj.LastName))
			Expect(userCustObj.State).To(Equal(userCustInputObj.State))
			Expect(userCustObj.Timezone).To(Equal(userCustInputObj.Timezone))
			Expect(userCustObj.UserStatus).To(Equal(userCustInputObj.UserStatus))
			Expect(userCustObj.Username).To(Equal(userCustInputObj.Username))
			Expect(permissions).To(ContainElement(userCustObj.Permissions[0].KeyName))
			Expect(permissions).To(ContainElement(userCustObj.Permissions[1].KeyName))
		})

	})

	Context("#GetObject", func() {
		BeforeEach(func() {
			userCustObj.Id = 631231
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err =
				test_helpers.ReadJsonTestFixtures("services", "SoftLayer_User_Customer_createAndGetObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves SoftLayer_User_Customer instance", func() {
			userCustObj, err := userCustomerService.GetObject(userCustObj.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(userCustObj.Id).ToNot(BeNil())
			Expect(userCustObj.ParentId).ToNot(BeNil())
			Expect(userCustObj.Address1).To(Equal(userCustInputObj.Address1))
			Expect(userCustObj.City).To(Equal(userCustInputObj.City))
			Expect(userCustObj.CompanyName).To(Equal(userCustInputObj.CompanyName))
			Expect(userCustObj.Country).To(Equal(userCustInputObj.Country))
			Expect(userCustObj.Email).To(Equal(userCustInputObj.Email))
			Expect(userCustObj.FirstName).To(Equal(userCustInputObj.FirstName))
			Expect(userCustObj.LastName).To(Equal(userCustInputObj.LastName))
			Expect(userCustObj.State).To(Equal(userCustInputObj.State))
			Expect(userCustObj.Timezone).To(Equal(userCustInputObj.Timezone))
			Expect(userCustObj.UserStatus).To(Equal(userCustInputObj.UserStatus))
			Expect(userCustObj.Username).To(Equal(userCustInputObj.Username))
			Expect(permissions).To(ContainElement(userCustObj.Permissions[0].KeyName))
			Expect(permissions).To(ContainElement(userCustObj.Permissions[1].KeyName))
		})
	})

	Context("#EditObject", func() {
		BeforeEach(func() {
			// Now prepare the fields to be updated.
			userCustInputObj = datatypes.SoftLayer_User_Customer{
				Address1:    "425 Market St",
				City:        "San Francisco",
				CompanyName: "TWC an IBM company",
				Country:     "US",
				Email:       "user4testingAgain@twc.ibm.com",
				FirstName:   "FirstName",
				LastName:    "LastName",
				Permissions: []datatypes.SoftLayer_User_Customer_CustomerPermission_Permission{
					datatypes.SoftLayer_User_Customer_CustomerPermission_Permission{
						KeyName: permissions[0],
					},
					datatypes.SoftLayer_User_Customer_CustomerPermission_Permission{
						KeyName: permissions[1],
					},
				},
				State:      "CA",
				Timezone:   110,
				UserStatus: 1001,
			}
			// Next prepare the mock response for the edit call.
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err =
				test_helpers.ReadJsonTestFixtures("services", "SoftLayer_User_Customer_editObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("edits a SoftLayer_User_Customer instance", func() {
			rv, err := userCustomerService.EditObject(1, userCustInputObj) // the userid value does not matter
			Expect(err).ToNot(HaveOccurred())
			Expect(rv).To(BeTrue())
		})
	})

	Context("#DeleteObject", func() {
		BeforeEach(func() {
			// Next prepare the mock response for the delete call.
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err =
				test_helpers.ReadJsonTestFixtures("services", "SoftLayer_User_Customer_deleteObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully deletes the SoftLayer_User_Customer instance", func() {
			// In SoftLayer delete User_Customer does not immediately delete the resource.
			// The delete api is really an 'editObject' api call, and which just marks
			// the userStatus field as "CANCEL_PENDING" (int id = 1021).
			// The actual record is deleted at some undetermined time in future by the SoftLayer
			// backend system.
			// While the user record is not completely deleted, one cannot create a new login
			// with the same username.
			rv, err := userCustomerService.DeleteObject(1) // the parameter value does not matter
			Expect(err).ToNot(HaveOccurred())
			Expect(rv).To(BeTrue())
		})
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

	Context("#RemoveApiAuthenticationKey", func() {
		BeforeEach(func() {
		})

		It("removes the api authentication key for a given softlayer user", func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err =
				test_helpers.ReadJsonTestFixtures("services", "SoftLayer_User_Customer_removeApiAuthenticationKey.json")
			Expect(err).ToNot(HaveOccurred())
			apiKeysToDelete := []datatypes.SoftLayer_User_Customer_ApiAuthentication{
				datatypes.SoftLayer_User_Customer_ApiAuthentication{
					Id:                1,
					AuthenticationKey: "456789",
					UserId:            45,
				},
			}
			var rv bool
			rv, err = userCustomerService.RemoveApiAuthenticationKey(45, apiKeysToDelete)
			Expect(err).ToNot(HaveOccurred())
			Expect(rv).To(BeTrue())
		})
	})

	Context("#GetPermissions", func() {
		BeforeEach(func() {
			userCustObj.Id = 631231
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err =
				test_helpers.ReadJsonTestFixtures("services", "SoftLayer_User_Customer_getPermissions.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves the permissions for a given SoftLayer_User_Customer instance", func() {
			var perms []datatypes.SoftLayer_User_Customer_CustomerPermission_Permission
			perms, err = userCustomerService.GetPermissions(userCustObj.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(permissions).To(ContainElement(perms[0].KeyName))
			Expect(permissions).To(ContainElement(perms[1].KeyName))
		})
	})

	Context("#AddBulkPortalPermission", func() {
		BeforeEach(func() {
			userCustObj.Id = 631231
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err =
				test_helpers.ReadJsonTestFixtures("services", "SoftLayer_User_Customer_addBulkPortalPermission.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully adds the given permissions to a given SoftLayer_User_Customer instance", func() {
			var perms = []datatypes.SoftLayer_User_Customer_CustomerPermission_Permission{
				datatypes.SoftLayer_User_Customer_CustomerPermission_Permission{
					KeyName: permissions[0],
				},
				datatypes.SoftLayer_User_Customer_CustomerPermission_Permission{
					KeyName: permissions[1],
				},
			}
			err = userCustomerService.AddBulkPortalPermission(userCustObj.Id, perms)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("#RemoveBulkPortalPermission", func() {
		BeforeEach(func() {
			userCustObj.Id = 631231
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err =
				test_helpers.ReadJsonTestFixtures("services", "SoftLayer_User_Customer_removeBulkPortalPermission.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully removes the specified permissions from a SoftLayer_User_Customer instance", func() {
			var perms = []datatypes.SoftLayer_User_Customer_CustomerPermission_Permission{
				datatypes.SoftLayer_User_Customer_CustomerPermission_Permission{
					KeyName: permissions[0],
				},
				datatypes.SoftLayer_User_Customer_CustomerPermission_Permission{
					KeyName: permissions[1],
				},
			}
			err = userCustomerService.RemoveBulkPortalPermission(userCustObj.Id, perms)
			Expect(err).ToNot(HaveOccurred())
		})
	})

})
