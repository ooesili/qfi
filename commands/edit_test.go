package commands_test

import (
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
		It("calls driver.Edit", func() {
			err := cmd.Run([]string{"foobar"})
			Expect(err).ToNot(HaveOccurred())
			Expect(driver.name).To(Equal("foobar"))
		})
	})

	Context("with more than one argument", func() {
		It("it returns an error", func() {
			err := cmd.Run([]string{"foobar", "bizbaz"})
			Expect(err).To(MatchError(ErrTooManyArgs))
		})
	})
})

type mockEditDriver struct {
	name string
}

func (d *mockEditDriver) Edit(name string) error {
	d.name = name
	return nil
}
