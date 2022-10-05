package thirteen_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestThirteen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Thirteen Suite")
}
