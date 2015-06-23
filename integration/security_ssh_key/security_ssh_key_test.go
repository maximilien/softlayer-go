package security_ssh_key_test

import (
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	softlayer "github.com/maximilien/softlayer-go/softlayer"
	testhelpers "github.com/maximilien/softlayer-go/test_helpers"
)

var _ = Describe("SoftLayer Security SSH keys", func() {
	var (
		err             error
		securityService softlayer.SoftLayer_Security_Ssh_Key_Service
	)

	BeforeEach(func() {
		securityService, err = testhelpers.CreateSecuritySshKeyService()
		Expect(err).ToNot(HaveOccurred())

		testhelpers.TIMEOUT = 30 * time.Second
		testhelpers.POLLING_INTERVAL = 10 * time.Second
	})

	Context("SoftLayer_Security_Ssh_Key", func() {
		It("creates an SSH key, update it, and delete it", func() {
			sshKeyPath := os.Getenv("SOFTLAYER_GO_TEST_SSH_KEY_PATH1")
			Expect(sshKeyPath).ToNot(Equal(""), "SOFTLAYER_GO_TEST_SSH_KEY_PATH1 env variable is not set")

			createdSshKey := testhelpers.CreateTestSshKey(sshKeyPath)
			testhelpers.WaitForCreatedSshKeyToBePresent(createdSshKey.Id)

			sshKeyService, err := testhelpers.CreateSecuritySshKeyService()
			Expect(err).ToNot(HaveOccurred())

			result, err := sshKeyService.GetObject(createdSshKey.Id)
			Expect(err).ToNot(HaveOccurred())

			Expect(result.CreateDate).ToNot(BeNil())
			Expect(result.Key).ToNot(Equal(""))
			Expect(result.Label).To(Equal("TEST:softlayer-go"))
			Expect(result.Notes).To(Equal("TEST:softlayer-go"))
			Expect(result.ModifyDate).To(BeNil())

			result.Label = "TEST:softlayer-go:edited-label"
			result.Notes = "TEST:softlayer-go:edited-notes"
			sshKeyService.EditObject(createdSshKey.Id, result)

			result2, err := sshKeyService.GetObject(createdSshKey.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(result2.Label).To(Equal("TEST:softlayer-go:edited-label"))
			Expect(result2.Notes).To(Equal("TEST:softlayer-go:edited-notes"))
			Expect(result2.ModifyDate).ToNot(BeNil())

			deleted, err := sshKeyService.DeleteObject(createdSshKey.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(deleted).To(BeTrue())

			testhelpers.WaitForDeletedSshKeyToNoLongerBePresent(createdSshKey.Id)
		})
	})
})
