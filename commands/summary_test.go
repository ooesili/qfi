package commands_test

import (
	"bytes"

	. "github.com/ooesili/qfi/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Summary", func() {
	var (
		cmd    Summary
		driver *mockSummaryDriver
		logger *bytes.Buffer
	)

	BeforeEach(func() {
		driver = &mockSummaryDriver{}
		logger = &bytes.Buffer{}
		cmd = Summary{driver, logger}
	})

	Context("with no arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{})
			Expect(err).ToNot(HaveOccurred())
			Expect(logger.String()).To(Equal("mockSummaryDriver.Summary was here"))
		})
	})

	Context("with more than zero arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{"foobar"})
			Expect(err).To(MatchError("too many arguments"))
		})
	})
})

type mockSummaryDriver struct{}

func (mockSummaryDriver) Summary() string {
	return "mockSummaryDriver.Summary was here"
}