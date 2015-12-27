package detect_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDetect(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Detect Suite")
}
