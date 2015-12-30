package edit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEdit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Edit Suite")
}
