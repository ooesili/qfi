package scripts_test

import (
	. "github.com/ooesili/qfi/scripts"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scripts.GetScript", func() {
	var scripts Scripts

	BeforeEach(func() {
		scripts = Scripts{}
	})

	Context("with a valid shell and script type", func() {
		It("finds an asset", func() {
			for _, shellName := range SupportedShells {
				for _, scriptType := range []string{"comp", "wrapper"} {
					data, err := scripts.GetScript(shellName, scriptType)
					Expect(err).ToNot(HaveOccurred(),
						"error occured finding %s_%s", shellName, scriptType)
					Expect(data).ToNot(BeNil())
				}
			}
		})
	})

	Context("when given an unsupported shell", func() {
		It("retuns an error", func() {
			_, err := scripts.GetScript("notashell", "wrapper")
			Expect(err).To(MatchError("unsupported shell: notashell"))
		})
	})

	Context("when given an invalid script type", func() {
		It("retuns an error", func() {
			_, err := scripts.GetScript("zsh", "thingdoer")
			Expect(err).To(MatchError("invalid script type: thingdoer"))
		})
	})
})
