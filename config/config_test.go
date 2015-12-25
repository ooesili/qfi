package config_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/ooesili/qfi/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var (
		configDir string
	)

	BeforeEach(func() {
		var err error
		configDir, err = ioutil.TempDir("", "qfi-test-")
		Expect(err).ToNot(HaveOccurred())

		// simple symlink helper function
		symlink := func(dir, name string) {
			err := os.Symlink(dir, filepath.Join(configDir, name))
			Expect(err).ToNot(HaveOccurred())
		}

		symlink("/foo/bar", "foobar")
		symlink("/biz/baz", "bizbaz")
		symlink("/foo/bar/qux", "qux")
	})

	AfterEach(func() {
		os.RemoveAll(configDir)
	})

	Describe("New", func() {
		Context("when the config directory containing non-symlink files", func() {
			It("returns an error", func() {
				badFile := filepath.Join(configDir, "notalink")
				file, err := os.Create(badFile)
				Expect(err).ToNot(HaveOccurred())
				file.Close()

				_, err = New(configDir)
				Expect(err).To(MatchError(
					fmt.Sprintf("not a symlink: %s", badFile),
				))
			})
		})

		Context("when a nonexistent directory is given", func() {
			It("returns an error", func() {
				badConfigDir := filepath.Join(configDir, "notadir")

				_, err := New(badConfigDir)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(HavePrefix(
					fmt.Sprintf("cannot open directory: %s: ", badConfigDir),
				))
			})
		})

		Context("when given a normal file is given", func() {
			It("returns an error", func() {
				badConfigDir := filepath.Join(configDir, "notadir")
				file, err := os.Create(badConfigDir)
				Expect(err).ToNot(HaveOccurred())
				file.Close()

				_, err = New(badConfigDir)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(HavePrefix(
					fmt.Sprintf("cannot read directory: %s: ", badConfigDir),
				))
			})
		})
	})

	Describe("Config.Resolve", func() {
		It("can resolve destinations of targets", func() {
			c, err := New(configDir)
			Expect(err).ToNot(HaveOccurred())
			Expect(c.Resolve("foobar")).To(Equal("/foo/bar"))
			Expect(c.Resolve("bizbaz")).To(Equal("/biz/baz"))
			Expect(c.Resolve("qux")).To(Equal("/foo/bar/qux"))
		})
	})

	Describe("Config.Add", func() {
		It("can add a target when given an relative path", func() {
			// add taraget
			c, err := New(configDir)
			Expect(err).NotTo(HaveOccurred())
			err = c.Add("newtarget", "/dingus")
			Expect(err).NotTo(HaveOccurred())

			// read the target's symlink
			destination, err := os.Readlink(filepath.Join(configDir, "newtarget"))
			Expect(err).NotTo(HaveOccurred())

			// make sure link points to the right place
			Expect(destination).To(Equal("/dingus"))
		})

		It("can add a target when given an relative path", func() {
			// add taraget
			c, err := New(configDir)
			Expect(err).NotTo(HaveOccurred())
			err = c.Add("newtarget", "dingus")
			Expect(err).NotTo(HaveOccurred())

			// read the target's symlink
			destination, err := os.Readlink(filepath.Join(configDir, "newtarget"))
			Expect(err).NotTo(HaveOccurred())

			// make sure link points to the right place
			pwd, err := os.Getwd()
			Expect(err).NotTo(HaveOccurred())
			Expect(destination).To(Equal(filepath.Join(pwd, "dingus")))
		})

		Context("when config directory is read only", func() {
			It("returns an error", func() {
				err := os.Chmod(configDir, 0555)
				Expect(err).NotTo(HaveOccurred())

				c, err := New(configDir)
				Expect(err).NotTo(HaveOccurred())

				err = c.Add("newtarget", "dingus")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(HavePrefix("cannot create symlink: %s: ",
					filepath.Join(configDir, "newtarget")))
			})
		})
	})

	Describe("Config.List", func() {
		It("lists all target names in alphabetical order", func() {
			c, err := New(configDir)
			Expect(err).ToNot(HaveOccurred())

			Expect(c.List()).To(Equal([]string{
				"bizbaz", "foobar", "qux",
			}))
		})
	})
})