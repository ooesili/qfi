package commands_test

import (
	"github.com/maraino/go-mock"
	. "github.com/ooesili/qfi/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add", func() {
	var (
		cmd    Add
		driver *mockAddDriver
	)

	BeforeEach(func() {
		driver = &mockAddDriver{}
		cmd = Add{driver}
	})

	Context("with no arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{})
			Expect(err).To(MatchError(ErrNoTargetOrFile))
		})
	})

	Context("with exactly one argument", func() {
		It("uses the basename of the argument as the target", func() {
			driver.When("Add", "baz", "/biz/baz")
			err := cmd.Run([]string{"/biz/baz"})
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("with exactly two arguments", func() {
		It("calls Driver.Add with the two arguments", func() {
			driver.When("Add", "foobar", "/foo/bar")
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

type mockAddDriver struct{ mock.Mock }

func (d *mockAddDriver) Add(name, destination string) error {
	ret := d.Called(name, destination)
	return ret.Error(1)
}
