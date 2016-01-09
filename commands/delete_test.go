package commands_test

import (
	"github.com/maraino/go-mock"
	. "github.com/ooesili/qfi/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete", func() {
	var (
		cmd    Delete
		driver *mockDeleteDriver
	)

	BeforeEach(func() {
		driver = &mockDeleteDriver{}
		cmd = Delete{driver}
	})

	Context("with no arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{})
			Expect(err).To(MatchError(ErrNoTargets))
		})
	})

	Context("with exactly one argument", func() {
		It("calls Driver.Delete", func() {
			driver.When("Delete", []string{"foobar"})
			err := cmd.Run([]string{"foobar"})
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("with exactly two arguments", func() {
		It("calls Driver.Delete", func() {
			driver.When("Delete", []string{"foobar", "bizbaz"})
			err := cmd.Run([]string{"foobar", "bizbaz"})
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("with more than two arguments", func() {
		It("calls Driver.Delete", func() {
			driver.When("Delete", []string{"foobar", "bizbaz", "qux", "boofar"})
			err := cmd.Run([]string{"foobar", "bizbaz", "qux", "boofar"})
			Expect(err).ToNot(HaveOccurred())
		})
	})
})

type mockDeleteDriver struct{ mock.Mock }

func (d *mockDeleteDriver) Delete(names ...string) error {
	ret := d.Called(names)
	return ret.Error(0)
}
