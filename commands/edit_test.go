package commands_test

import (
	"github.com/maraino/go-mock"
	. "github.com/ooesili/qfi/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Edit", func() {
	var (
		cmd    Edit
		driver *mockEditDriver
	)

	BeforeEach(func() {
		driver = &mockEditDriver{}
		cmd = Edit{driver}
	})

	Context("with no arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{})
			Expect(err).To(MatchError(ErrNoTarget))
		})
	})

	Context("with one argument", func() {
		It("calls Driver.Edit", func() {
			driver.When("Edit", "foobar")
			err := cmd.Run([]string{"foobar"})
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("with more than one argument", func() {
		It("it returns an error", func() {
			err := cmd.Run([]string{"foobar", "bizbaz"})
			Expect(err).To(MatchError(ErrTooManyArgs))
		})
	})
})

type mockEditDriver struct{ mock.Mock }

func (d *mockEditDriver) Edit(name string) error {
	ret := d.Called(name)
	return ret.Error(0)
}
