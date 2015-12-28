package detect_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/ooesili/qfi/detect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Detect", func() {
	var (
		tempDir  string
		detector Detector
	)

	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "qfi-test-")
		Expect(err).ToNot(HaveOccurred())
		detector = Detector{}
	})

	AfterEach(func() {
		err := os.RemoveAll(tempDir)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("when given a normal file", func() {
		It("returns NormalFile", func() {
			file := filepath.Join(tempDir, "foobar")
			fi, err := os.Create(file)
			Expect(err).ToNot(HaveOccurred())
			fi.Close()

			Expect(detector.Detect(file)).To(Equal(NormalFile))
		})
	})

	Context("when given a file without write permissions", func() {
		It("returns UnwritableFile", func() {
			file := filepath.Join(tempDir, "foobar")
			fi, err := os.Create(file)
			Expect(err).ToNot(HaveOccurred())
			fi.Chmod(0400)
			Expect(err).ToNot(HaveOccurred())
			fi.Close()

			Expect(detector.Detect(file)).To(Equal(UnwritableFile))
		})
	})

	Context("when given a file inside of an unreadable directory", func() {
		BeforeEach(func() {
			err := os.Chmod(tempDir, 0000)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			err := os.Chmod(tempDir, 0700)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns InaccessibleFile", func() {
			file := filepath.Join(tempDir, "foobar")
			Expect(detector.Detect(file)).To(Equal(InaccessibleFile))
		})
	})

	Context("when given a nonexistent file", func() {
		It("returns NonexistentFile", func() {
			file := filepath.Join(tempDir, "foobar")
			Expect(detector.Detect(file)).To(Equal(NonexistentFile))
		})
	})

	Context("when given a readable directory", func() {
		It("returns NormalDirectory", func() {
			file := filepath.Join(tempDir, "foobar")
			err := os.Mkdir(file, 0700)
			Expect(err).ToNot(HaveOccurred())

			Expect(detector.Detect(file)).To(Equal(NormalDirectory))
		})
	})

	Context("when given an unreadable directory", func() {
		It("returns UnreadableDirectory", func() {
			file := filepath.Join(tempDir, "foobar")
			err := os.Mkdir(file, 0000)
			Expect(err).ToNot(HaveOccurred())

			Expect(detector.Detect(file)).To(Equal(UnreadableDirectory))
		})
	})

	Context("when a given path's parent directory is actually a file", func() {
		It("returns UnknownFile", func() {
			badParent := filepath.Join(tempDir, "badparent")
			fi, err := os.Create(badParent)
			Expect(err).ToNot(HaveOccurred())
			fi.Close()

			file := filepath.Join(badParent, "foobar")
			Expect(detector.Detect(file)).To(Equal(UnknownFile))
		})
	})
})
