package commands_test

import (
	"github.com/maraino/go-mock"
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
			Expect(err).To(MatchError(ErrNoTarget))
		})
	})

	Context("with exactly one argument", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{"foobar"})
			Expect(err).To(MatchError(ErrNoFile))
		})
	})

	Context("with exactly two arguments", func() {
		It("calls Driver.Move", func() {
			driver.When("Move", "foobar", "/foo/bar")
			err := cmd.Run([]string{"foobar", "/foo/bar"})
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("with more than two arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{"foobar", "/foo/bar", "oops"})
			Expect(err).To(MatchError(ErrTooManyArgs))
		})
	})
})

type mockMoveDriver struct{ mock.Mock }

func (d *mockMoveDriver) Move(name, destination string) error {
	ret := d.Called(name, destination)
	return ret.Error(0)
}
