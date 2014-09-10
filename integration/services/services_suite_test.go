package services_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	testhelpers "github.com/maximilien/softlayer-go/test_helpers"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration: Services Suite")
}

func cleanUpTestResources() {
	virtualGuestIds, err := testhelpers.FindAndDeleteTestVirtualGuests()
	Expect(err).ToNot(HaveOccurred())

	for _, vgId := range virtualGuestIds {
		waitForVirtualGuestToHaveNoActiveTransactions(vgId)
	}

	err = testhelpers.FindAndDeleteTestSshKeys()
	Expect(err).ToNot(HaveOccurred())
}
