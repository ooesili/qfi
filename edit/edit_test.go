package edit_test

import (
	"github.com/maraino/go-mock"
	"github.com/ooesili/qfi/detect"
	. "github.com/ooesili/qfi/edit"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Edit", func() {
	Context("when given a normal file", func() {
		It("calls the detector and resolver with the right args", func() {
			resolver := &mockResolver{}
			detector := &mockDetector{}
			executor := &mockExecutor{}
			resolver.When("Resolve", "foobar").Return("/foo/bar")
			detector.When("Detect", "/foo/bar").Return(detect.NormalFile)
			executor.When("Exec", "vim", []string{"/foo/bar"})

			editor := Editor{"vim", detector, executor, resolver}
			err := editor.Edit("foobar")
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when the resolver returns a specific file type", func() {
		It("executes the correct editor", func() {
			destination := "/biz/baz"
			tests := []struct {
				fileType    detect.Type
				typeString  string
				commandName string
				commandArgs []string
			}{
				{detect.NormalFile, "NormalFile",
					"emacs", []string{destination}},
				{detect.UnwritableFile, "UnwritableFile",
					"sudo", []string{"-e", destination}},
				{detect.InaccessibleFile, "InaccessibleFile",
					"sudo", []string{"-e", destination}},
				{detect.NonexistentFile, "NonexistentFile",
					"emacs", []string{destination}},
			}

			for _, test := range tests {
				detector := &mockDetector{}
				executor := &mockExecutor{}
				resolver := &mockResolver{}
				resolver.When("Resolve", "foobar").Return(destination)
				detector.When("Detect", destination).Return(test.fileType)
				executor.When("Exec", test.commandName, test.commandArgs)

				editor := Editor{"emacs", detector, executor, resolver}
				err := editor.Edit("foobar")
				Expect(err).ToNot(HaveOccurred(),
					"Edit should not fail for %s", test.typeString)
			}
		})
	})

	Context("when the resolver returns a directory type", func() {
		It("should return ErrWrapperShouldChdir", func() {
			tests := []struct {
				fileType   detect.Type
				typeString string
			}{
				{detect.NormalDirectory, "NormalDirectory"},
				{detect.UnreadableDirectory, "UnreadableDirectory"},
			}

			for _, test := range tests {
				detector := &mockDetector{}
				executor := &mockExecutor{}
				resolver := &mockResolver{}
				resolver.When("Resolve", "foobar").Return("/foo/bar")
				detector.When("Detect", "/foo/bar").Return(test.fileType)

				editor := Editor{"vim", detector, executor, resolver}
				err := editor.Edit("foobar")
				Expect(err).To(MatchError(ErrWrapperShouldChdir),
					"Edit return error for %s", test.typeString)
			}
		})
	})

	Context("when the resolver returns UnknownFile", func() {
		It("should return an error", func() {
			detector := &mockDetector{}
			executor := &mockExecutor{}
			resolver := &mockResolver{}
			resolver.When("Resolve", "foobar").Return("/foo/bar")
			detector.When("Detect", "/foo/bar").Return(detect.UnknownFile)

			editor := Editor{"vim", detector, executor, resolver}
			err := editor.Edit("foobar")
			Expect(err).To(MatchError("unknown file type for: /foo/bar"))
		})
	})
})

type mockDetector struct{ mock.Mock }

func (d *mockDetector) Detect(path string) detect.Type {
	ret := d.Called(path)
	return ret.Get(0).(detect.Type)
}

type mockResolver struct{ mock.Mock }

func (r *mockResolver) Resolve(name string) (string, error) {
	ret := r.Called(name)
	return ret.String(0), ret.Error(1)
}

type mockExecutor struct{ mock.Mock }

func (e *mockExecutor) Exec(name string, args ...string) error {
	ret := e.Called(name, args)
	return ret.Error(0)
}
