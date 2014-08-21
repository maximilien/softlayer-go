package data_types_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDataTypes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DataTypes Suite")
}
