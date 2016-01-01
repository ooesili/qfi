package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var qfiCmd string

func TestQfi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Qfi Suite")
}

var _ = BeforeSuite(func() {
	var err error
	qfiCmd, err = gexec.Build("github.com/ooesili/qfi")
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
