package commands_test

import (
	"bytes"

	"github.com/maraino/go-mock"
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
			driver.When("Summary").Return("mockSummaryDriver.Summary was here")
			err := cmd.Run([]string{})
			Expect(err).ToNot(HaveOccurred())
			Expect(logger.String()).To(Equal("mockSummaryDriver.Summary was here"))
		})
	})

	Context("with more than zero arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{"foobar"})
			Expect(err).To(MatchError(ErrTooManyArgs))
		})
	})
})

type mockSummaryDriver struct{ mock.Mock }

func (d *mockSummaryDriver) Summary() string {
	ret := d.Called()
	return ret.String(0)
}
