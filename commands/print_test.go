package commands_test

import (
	"bytes"

	. "github.com/ooesili/qfi/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Print", func() {
	var (
		cmd    Print
		logger *bytes.Buffer
	)

	BeforeEach(func() {
		logger = &bytes.Buffer{}
		cmd = Print{"this is some text", logger}
	})

	Context("with no arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{})
			Expect(err).ToNot(HaveOccurred())
			Expect(logger.String()).To(Equal("this is some text"))
		})
	})

	Context("with more than zero arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{"foobar"})
			Expect(err).To(MatchError(ErrTooManyArgs))
		})
	})
})
