package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestQfi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Qfi Suite")
}
