package commands_test

import (
	"bytes"

	"github.com/maraino/go-mock"
	. "github.com/ooesili/qfi/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shell", func() {
	var (
		cmd    Shell
		driver *mockShellDriver
		logger *bytes.Buffer
	)

	BeforeEach(func() {
		driver = &mockShellDriver{}
		logger = &bytes.Buffer{}
		cmd = Shell{driver, logger}
	})

	Context("with no arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{})
			Expect(err).To(MatchError(ErrNoShell))
		})
	})

	Context("with exactly one argument", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{"foobar"})
			Expect(err).To(MatchError(ErrNoScript))
		})
	})

	Context("with exactly two arguments", func() {
		It("calls Driver.GetScript and prints the results", func() {
			driver.When("GetScript", "zsh", "comp").
				Return([]byte("cool script, brah\n"))
			err := cmd.Run([]string{"zsh", "comp"})
			Expect(err).ToNot(HaveOccurred())
			Expect(logger.String()).To(Equal("cool script, brah\n"))
		})
	})

	Context("with more than two arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{"zsh", "comp", "oops"})
			Expect(err).To(MatchError(ErrTooManyArgs))
		})
	})
})

type mockShellDriver struct{ mock.Mock }

func (d *mockShellDriver) GetScript(shell, scriptType string) ([]byte, error) {
	ret := d.Called(shell, scriptType)
	return ret.Bytes(0), ret.Error(1)
}
