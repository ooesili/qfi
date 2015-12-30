package edit_test

import (
	"github.com/ooesili/qfi/detect"
	. "github.com/ooesili/qfi/edit"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Edit", func() {
	Context("when given a normal file", func() {
		It("calls the detector and resolver with the right args", func() {
			detector := &mockDetector{fileType: detect.NormalFile}
			executor := &mockExecutor{}
			resolver := &mockResolver{destination: "/foo/bar"}

			editor := Editor{"vim", detector, executor, resolver}
			err := editor.Edit("foobar")
			Expect(err).ToNot(HaveOccurred())

			Expect(resolver.calledName).To(Equal("foobar"))
			Expect(detector.calledPath).To(Equal("/foo/bar"))
			Expect(executor.calledName).To(Equal("vim"))
			Expect(executor.calledArgs).To(Equal([]string{"/foo/bar"}))
		})
	})

	Context("when the resolver returns a file type", func() {
		It("executes the correct editor", func() {
			tests := []struct {
				fileType    detect.Type
				typeString  string
				commandName string
			}{
				{detect.NormalFile, "NormalFile", "vim"},
				{detect.UnwritableFile, "UnwritableFile", "sudoedit"},
				{detect.InaccessibleFile, "InaccessibleFile", "sudoedit"},
				{detect.NonexistentFile, "NonexistentFile", "vim"},
			}

			for _, test := range tests {
				detector := &mockDetector{fileType: test.fileType}
				executor := &mockExecutor{}
				resolver := &mockResolver{destination: "/foo/bar"}

				editor := Editor{"vim", detector, executor, resolver}
				err := editor.Edit("foobar")
				Expect(err).ToNot(HaveOccurred(),
					"Edit should not fail for %s", test.typeString)

				Expect(executor.calledName).To(Equal(test.commandName),
					"should run %s when given %s", test.commandName, test.typeString)
			}
		})
	})

	Context("when the resolver returns a directory type", func() {
		It("should return an WrapperShouldChdirError", func() {
			tests := []struct {
				fileType   detect.Type
				typeString string
			}{
				{detect.NormalDirectory, "NormalDirectory"},
				{detect.UnreadableDirectory, "UnreadableDirectory"},
			}

			for _, test := range tests {
				detector := &mockDetector{fileType: test.fileType}
				executor := &mockExecutor{}
				resolver := &mockResolver{destination: "/foo/bar"}

				editor := Editor{"vim", detector, executor, resolver}
				err := editor.Edit("foobar")
				Expect(err).To(MatchError(WrapperShouldChdirError{"/foo/bar"}),
					"Edit return error for %s", test.typeString)
			}
		})
	})

	Context("when the resolver returns UnknownFile", func() {
		It("should return an error", func() {
			detector := &mockDetector{fileType: detect.UnknownFile}
			executor := &mockExecutor{}
			resolver := &mockResolver{destination: "/foo/bar"}

			editor := Editor{"vim", detector, executor, resolver}
			err := editor.Edit("foobar")
			Expect(err).To(MatchError("unknown file type for: /foo/bar"))
		})
	})
})

type mockDetector struct {
	calledPath string
	fileType   detect.Type
}

func (d *mockDetector) Detect(path string) detect.Type {
	d.calledPath = path
	return d.fileType
}

type mockResolver struct {
	calledName  string
	destination string
}

func (d *mockResolver) Resolve(name string) (string, error) {
	d.calledName = name
	return d.destination, nil
}

type mockExecutor struct {
	calledName string
	calledArgs []string
}

func (e *mockExecutor) Exec(name string, args ...string) error {
	e.calledName = name
	e.calledArgs = args
	return nil
}
