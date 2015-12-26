package commands_test

import (
	"bytes"
	"fmt"

	. "github.com/ooesili/qfi/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("List", func() {
	var (
		cmd    List
		driver *mockListDriver
		logger *bytes.Buffer
	)

	BeforeEach(func() {
		driver = &mockListDriver{
			targets: []string{"bizbaz", "foobar", "qux"},
		}
		logger = &bytes.Buffer{}
		cmd = List{driver, logger}
	})

	Context("with no arguments", func() {
		It("calls driver.List and prints the results", func() {
			err := cmd.Run([]string{})
			Expect(err).ToNot(HaveOccurred())
			Expect(logger.String()).To(Equal("bizbaz\nfoobar\nqux\n"))
		})
	})

	Context("with one argument", func() {
		It("calls driver.Resolve and prints the results", func() {
			err := cmd.Run([]string{"foobar"})
			Expect(err).ToNot(HaveOccurred())
			Expect(logger.String()).To(Equal("/path/to/foobar\n"))
		})
	})

	Context("with more than one argument", func() {
		It("it returns an error", func() {
			err := cmd.Run([]string{"foobar", "bizbaz"})
			Expect(err).To(MatchError("too many arguments"))
		})
	})
})

type mockListDriver struct {
	targets []string
}

func (d *mockListDriver) List() []string {
	return d.targets
}

func (d *mockListDriver) Resolve(name string) (string, error) {
	return fmt.Sprintf("/path/to/%s", name), nil
}
