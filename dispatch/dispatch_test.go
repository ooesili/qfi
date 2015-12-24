package dispatch_test

import (
	. "github.com/ooesili/qfi/dispatch"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dispatch", func() {
	var (
		calledArgs []string
		d          *Dispatcher
	)

	testCommand := CommandFunc(func(args []string) error {
		calledArgs = args
		return nil
	})

	BeforeEach(func() {
		d = New()
		calledArgs = nil
	})

	normalUsage := func() {
		d.Register("-s", testCommand)

		err := d.Run([]string{"-s", "arg1", "arg2"})
		Expect(err).NotTo(HaveOccurred())
		Expect(calledArgs).To(Equal([]string{"arg1", "arg2"}))
	}

	Context("without a fallback command", func() {
		It("runs a command with the correct args", normalUsage)

		It("returns an error when the command is not found", func() {
			err := d.Run([]string{"frooble"})
			Expect(err).To(Equal(ErrNotFound))
		})

		It("returns an error when no args are given", func() {
			err := d.Run([]string{})
			Expect(err).To(Equal(ErrNoArgs))
		})
	})

	Context("with a fallback command", func() {
		BeforeEach(func() {
			d.RegisterFallback(testCommand)
		})

		It("runs a command with the correct args", normalUsage)

		Context("when no registered command matches", func() {
			It("calls the fallback command with all of the given args", func() {
				err := d.Run([]string{"frooble", "grooble"})
				Expect(err).NotTo(HaveOccurred())
				Expect(calledArgs).To(Equal([]string{"frooble", "grooble"}))
			})
		})

		Context("when no arguments are given", func() {
			It("calls the fallback command with all of the given args", func() {
				err := d.Run([]string{})
				Expect(err).NotTo(HaveOccurred())
				Expect(calledArgs).To(Equal([]string{}))
			})
		})
	})
})
