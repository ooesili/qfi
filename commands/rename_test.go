package commands_test

import (
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
			Expect(err).To(MatchError("no target specified"))
		})
	})

	Context("with exactly one argument", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{"foobar"})
			Expect(err).To(MatchError("no file specified"))
		})
	})

	Context("with exactly two arguments", func() {
		It("calls Driver.Rename", func() {
			err := cmd.Run([]string{"foobar", "boofar"})
			Expect(err).ToNot(HaveOccurred())
			Expect(driver.name).To(Equal("foobar"))
			Expect(driver.newName).To(Equal("boofar"))
		})
	})

	Context("with more than two arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{"foobar", "boofar", "oops"})
			Expect(err).To(MatchError("too many arguments"))
		})
	})
})

type mockRenameDriver struct {
	name    string
	newName string
}

func (d *mockRenameDriver) Rename(name, newName string) error {
	d.name = name
	d.newName = newName
	return nil
}
