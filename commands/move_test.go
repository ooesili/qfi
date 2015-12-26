package commands_test

import (
	. "github.com/ooesili/qfi/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Move", func() {
	var (
		cmd    Move
		driver *mockMoveDriver
	)

	BeforeEach(func() {
		driver = &mockMoveDriver{}
		cmd = Move{driver}
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
		It("calls Driver.Move", func() {
			err := cmd.Run([]string{"foobar", "/foo/bar"})
			Expect(err).ToNot(HaveOccurred())
			Expect(driver.name).To(Equal("foobar"))
			Expect(driver.destination).To(Equal("/foo/bar"))
		})
	})

	Context("with more than two arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{"foobar", "/foo/bar", "oops"})
			Expect(err).To(MatchError("too many arguments"))
		})
	})
})

type mockMoveDriver struct {
	name        string
	destination string
}

func (d *mockMoveDriver) Move(name, destination string) error {
	d.name = name
	d.destination = destination
	return nil
}
