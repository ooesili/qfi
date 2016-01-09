package commands_test

import (
	"github.com/maraino/go-mock"
	. "github.com/ooesili/qfi/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rename", func() {
	var (
		cmd    Rename
		driver *mockRenameDriver
	)

	BeforeEach(func() {
		driver = &mockRenameDriver{}
		cmd = Rename{driver}
	})

	Context("with no arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{})
			Expect(err).To(MatchError(ErrNoTarget))
		})
	})

	Context("with exactly one argument", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{"foobar"})
			Expect(err).To(MatchError(ErrNoNewName))
		})
	})

	Context("with exactly two arguments", func() {
		It("calls Driver.Rename", func() {
			driver.When("Rename", "foobar", "boofar")
			err := cmd.Run([]string{"foobar", "boofar"})
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("with more than two arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{"foobar", "boofar", "oops"})
			Expect(err).To(MatchError(ErrTooManyArgs))
		})
	})
})

type mockRenameDriver struct{ mock.Mock }

func (d *mockRenameDriver) Rename(name, newName string) error {
	ret := d.Called(name, newName)
	return ret.Error(0)
}
