package commands_test

import (
	"bytes"

	"github.com/maraino/go-mock"
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
		driver = &mockListDriver{}
		logger = &bytes.Buffer{}
		cmd = List{driver, logger}
	})

	Context("with no arguments", func() {
		It("calls driver.List and prints the results", func() {
			driver.When("List").Return([]string{"bizbaz", "foobar", "qux"})
			err := cmd.Run([]string{})
			Expect(err).ToNot(HaveOccurred())
			Expect(logger.String()).To(Equal("bizbaz\nfoobar\nqux\n"))
		})
	})

	Context("with one argument", func() {
		It("calls driver.Resolve and prints the results", func() {
			driver.When("Resolve", "foobar").Return("/foo/bar")
			err := cmd.Run([]string{"foobar"})
			Expect(err).ToNot(HaveOccurred())
			Expect(logger.String()).To(Equal("/foo/bar\n"))
		})
	})

	Context("with more than one argument", func() {
		It("it returns an error", func() {
			err := cmd.Run([]string{"foobar", "bizbaz"})
			Expect(err).To(MatchError(ErrTooManyArgs))
		})
	})
})

type mockListDriver struct{ mock.Mock }

func (d *mockListDriver) List() []string {
	ret := d.Called()
	return ret.Get(0).([]string)
}

func (d *mockListDriver) Resolve(name string) (string, error) {
	ret := d.Called(name)
	return ret.String(0), ret.Error(1)
}
