package summarize_test

import (
	"github.com/fatih/color"
	"github.com/maraino/go-mock"
	"github.com/ooesili/qfi/detect"
	. "github.com/ooesili/qfi/summarize"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Summarize", func() {
	BeforeEach(func() {
		// turn off color before each test
		color.NoColor = true
	})

	Context("when given a single target", func() {
		It("displays it's summary with the right color and arrow", func() {
			color.NoColor = false

			// setup test cases
			tests := []struct {
				fileType   detect.Type
				typeString string
				arrowChar  rune
				fgColor    color.Attribute
			}{
				{detect.NormalFile, "NormalFile", '-', color.FgGreen},
				{detect.UnwritableFile, "UnwritableFile", '#', color.FgYellow},
				{detect.InaccessibleFile, "InaccessibleFile", '#', color.FgRed},
				{detect.NonexistentFile, "NonexistentFile", '-', color.FgRed},
				{detect.NormalDirectory, "NormalDirectory", '/', color.FgBlue},
				{detect.UnreadableDirectory, "UnreadableDirectory", '/', color.FgRed},
				{detect.UnknownFile, "UnknownFile", '?', color.FgMagenta},
			}

			// run tests
			for _, test := range tests {
				resolver := &mockResolver{}
				detector := &mockDetector{}
				resolver.When("List").Return([]string{"foobar"})
				resolver.When("Resolve", "foobar").Return("/foo/bar")
				detector.When("Detect", "/foo/bar").Return(test.fileType)

				colorSprintf := color.New(test.fgColor).SprintfFunc()

				summarizer := Summarizer{detector, resolver}
				summary := summarizer.Summary()
				Expect(summary).To(Equal(
					colorSprintf("foobar %c> /foo/bar", test.arrowChar)+"\n"),
					"when given detect.%s", test.typeString)
			}
		})
	})

	Context("when given multiple targets", func() {
		It("aligns the output into columns", func() {
			resolver := &mockResolver{}
			detector := &mockDetector{}
			resolver.When("List").Return([]string{"longer", "longest", "short"})
			resolver.When("Resolve", "longer").Return("/biz/baz")
			resolver.When("Resolve", "longest").Return("/foo/bar/qux")
			resolver.When("Resolve", "short").Return("/foo/bar")
			detector.When("Detect", "/biz/baz").Return(detect.UnwritableFile)
			detector.When("Detect", "/foo/bar/qux").Return(detect.NormalFile)
			detector.When("Detect", "/foo/bar").Return(detect.NormalDirectory)

			summarizer := Summarizer{detector, resolver}
			summary := summarizer.Summary()
			Expect(summary).To(Equal(`
longer  #> /biz/baz
longest -> /foo/bar/qux
short   /> /foo/bar
`[1:],
			))
		})
	})
})

type mockResolver struct{ mock.Mock }

func (r *mockResolver) List() []string {
	ret := r.Called()
	return ret.Get(0).([]string)
}

func (r *mockResolver) Resolve(name string) (string, error) {
	ret := r.Called(name)
	return ret.String(0), ret.Error(1)
}

type mockDetector struct{ mock.Mock }

func (d *mockDetector) Detect(path string) detect.Type {
	ret := d.Called(path)
	return ret.Get(0).(detect.Type)
}
