package commands_test

import (
	. "github.com/ooesili/qfi/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete", func() {
	var (
		cmd    Delete
		driver *mockDeleteDriver
	)

	BeforeEach(func() {
		driver = &mockDeleteDriver{}
		cmd = Delete{driver}
	})

	Context("with no arguments", func() {
		It("returns an error", func() {
			err := cmd.Run([]string{})
			Expect(err).To(MatchError(ErrNoTargets))
		})
	})

	Context("with exactly one argument", func() {
		It("calls Config.Delete", func() {
			err := cmd.Run([]string{"foobar"})
			Expect(err).ToNot(HaveOccurred())
			Expect(driver.names).To(Equal([]string{"foobar"}))
		})
	})

	Context("with exactly two arguments", func() {
		It("calls Driver.Delete", func() {
			err := cmd.Run([]string{"foobar", "bizbaz"})
			Expect(err).ToNot(HaveOccurred())
			Expect(driver.names).To(Equal([]string{"foobar", "bizbaz"}))
		})
	})

	Context("with more than two arguments", func() {
		It("calls Driver.Delete", func() {
			err := cmd.Run([]string{"foobar", "bizbaz", "qux", "boofar"})
			Expect(err).ToNot(HaveOccurred())
			Expect(driver.names).To(Equal(
				[]string{"foobar", "bizbaz", "qux", "boofar"}))
		})
	})
})

type mockDeleteDriver struct {
	names []string
}

func (d *mockDeleteDriver) Delete(names ...string) error {
	d.names = names
	return nil
}
