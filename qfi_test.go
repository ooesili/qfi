package main_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Qfi", func() {
	var (
		tempDir   string
		configDir string

		somefile  string
		otherfile string
		somedir   string
		extrafile string
	)

	BeforeEach(func() {
		var err error
		// create temp dir
		tempDir, err = ioutil.TempDir("", "qfi-test-")
		Expect(err).ToNot(HaveOccurred())

		// make sure a config directory is automatically created
		configDir = filepath.Join(tempDir, "config")

		// set QFI_CONFIGDIR
		err = os.Setenv("QFI_CONFIGDIR", configDir)
		Expect(err).ToNot(HaveOccurred())

		// create some files
		touch := func(path string) {
			fi, err := os.Create(path)
			Expect(err).ToNot(HaveOccurred())
			fi.Close()
		}
		somefile = filepath.Join(tempDir, "somefile")
		touch(somefile)
		otherfile = filepath.Join(tempDir, "otherfile")
		touch(otherfile)
		extrafile = filepath.Join(tempDir, "extrafile")
		touch(extrafile)

		// create a directory
		somedir = filepath.Join(tempDir, "somedir")
		err = os.Mkdir(somedir, 0700)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		err := os.RemoveAll(configDir)
		Expect(err).ToNot(HaveOccurred())
	})

	qfi := func(args ...string) (string, string, error) {
		out, err := exec.Command(
			qfiCmd, args...,
		).CombinedOutput()
		desc := fmt.Sprintf("$ qfi %s\n%s", strings.Join(args, " "), out)
		return string(out), desc, err
	}

	Context("after adding some targets", func() {
		BeforeEach(func() {
			_, desc, err := qfi("-a", "some", somefile)
			Expect(err).ToNot(HaveOccurred(), desc)
			_, desc, err = qfi("-a", "other", otherfile)
			Expect(err).ToNot(HaveOccurred(), desc)
			_, desc, err = qfi("-a", "dir", somedir)
			Expect(err).ToNot(HaveOccurred(), desc)
		})

		It("can move targets", func() {
			out, desc, err := qfi("-m", "some", extrafile)
			Expect(err).ToNot(HaveOccurred(), desc)

			out, desc, err = qfi("-l", "some")
			Expect(err).ToNot(HaveOccurred(), desc)
			Expect(out).To(Equal(extrafile + "\n"))
		})

		It("can list targets", func() {
			out, desc, err := qfi("-l")
			Expect(err).ToNot(HaveOccurred(), desc)
			Expect(out).To(Equal("dir\nother\nsome\n"))
		})

		It("can rename targets", func() {
			out, desc, err := qfi("-r", "some", "newname")
			Expect(err).ToNot(HaveOccurred(), desc)

			out, desc, err = qfi("-l")
			Expect(err).ToNot(HaveOccurred(), desc)
			Expect(out).To(Equal("dir\nnewname\nother\n"))
		})

		It("can delete targets", func() {
			out, desc, err := qfi("-d", "dir", "other")
			Expect(err).ToNot(HaveOccurred(), desc)

			out, desc, err = qfi("-l")
			Expect(err).ToNot(HaveOccurred(), desc)
			Expect(out).To(Equal("some\n"))
		})

		It("can summarize targets", func() {
			out, desc, err := qfi("-s")
			Expect(err).ToNot(HaveOccurred(), desc)
			Expect(out).To(Equal(fmt.Sprintf(`
dir   /> %s
other -> %s
some  -> %s
`[1:], somedir, otherfile, somefile)))
		})

		It("can edit a target that points to a file", func() {
			err := os.Setenv("EDITOR", "echo")
			Expect(err).ToNot(HaveOccurred())

			out, desc, err := qfi("some")
			Expect(err).ToNot(HaveOccurred(), desc)
			Expect(out).To(Equal(somefile + "\n"))
		})

		It("exits with 2 when given a target that points a directory", func() {
			err := os.Setenv("EDITOR", "echo")
			Expect(err).ToNot(HaveOccurred())

			_, _, err = qfi("dir")
			Expect(err).To(MatchError("exit status 2"))
		})
	})

	It("can print the version without returning an error", func() {
		_, desc, err := qfi("--version")
		Expect(err).ToNot(HaveOccurred(), desc)
	})

	It("can print the help message without returning an error", func() {
		_, desc, err := qfi("--help")
		Expect(err).ToNot(HaveOccurred(), desc)
	})
})
