package commands_test

import (
	"reflect"

	. "github.com/ooesili/qfi/commands"
	"github.com/ooesili/qfi/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Commands", func() {
	var (
		cmds Commands
		cfg  *mockConfig
	)

	BeforeEach(func() {
		cfg = &mockConfig{}
		cmds = New(cfg)
	})

	Describe("Config", func() {
		It("is implemented by config.Config", func() {
			realType := reflect.TypeOf(config.Config{})
			interfaceType := reflect.TypeOf((*Config)(nil)).Elem()

			Expect(realType.Implements(interfaceType)).To(BeTrue())
		})
	})

	Describe("Add", func() {
		Context("with no arguments", func() {
			It("returns an error", func() {
				err := cmds.Add([]string{})
				Expect(err).To(MatchError("no target specified"))
			})
		})

		Context("with exactly one argument", func() {
			It("returns an error", func() {
				err := cmds.Add([]string{"foobar"})
				Expect(err).To(MatchError("no file specified"))
			})
		})

		Context("with exactly two arguments", func() {
			It("calls Config.Add", func() {
				err := cmds.Add([]string{"foobar", "/foo/bar"})
				Expect(err).ToNot(HaveOccurred())
				Expect(cfg.addArgs.name).To(Equal("foobar"))
				Expect(cfg.addArgs.destination).To(Equal("/foo/bar"))
			})
		})

		Context("with more than two arguments", func() {
			It("returns an error", func() {
				err := cmds.Add([]string{"foobar", "/foo/bar", "oops"})
				Expect(err).To(MatchError("too many arguments"))
			})
		})
	})
})

type mockConfig struct {
	addArgs struct{ name, destination string }
}

func (c *mockConfig) Add(name, destination string) error {
	c.addArgs.name = name
	c.addArgs.destination = destination
	return nil
}
