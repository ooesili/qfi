package edit_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/ooesili/qfi/edit"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RealExecutor", func() {
	var tempDir string

	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "qfi-test-")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		err := os.RemoveAll(tempDir)
		Expect(err).ToNot(HaveOccurred())
	})

	It("calls a real shell command", func() {
		touchFile := filepath.Join(tempDir, "foobar")

		executor := RealExecutor{}
		err := executor.Exec("touch", touchFile)

		_, err = os.Stat(touchFile)
		Expect(err).ToNot(HaveOccurred())
	})
})
