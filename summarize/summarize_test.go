package summarize_test

import (
	"github.com/fatih/color"
	"github.com/ooesili/qfi/detect"
	. "github.com/ooesili/qfi/summarize"

	"sort"

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
				resolver := mockResolver(map[string]string{
					"foobar": "/foo/bar",
				})
				detector := mockDetector(map[string]detect.Type{
					"/foo/bar": test.fileType,
				})

				colorSprintf := color.New(test.fgColor).SprintfFunc()

				summarizer := Summarizer{detector, resolver}
				summary := summarizer.Summary()
				Expect(summary).To(Equal(
					colorSprintf("foobar %c> /foo/bar\n", test.arrowChar)),
					"when given detect.%s", test.typeString)
			}
		})
	})

	Context("when given multiple targets", func() {
		It("aligns the output into columns", func() {
			resolver := mockResolver(map[string]string{
				"short":   "/foo/bar",
				"longer":  "/biz/baz",
				"longest": "/foo/bar/qux",
			})
			detector := mockDetector(map[string]detect.Type{
				"/foo/bar":     detect.NormalDirectory,
				"/biz/baz":     detect.UnwritableFile,
				"/foo/bar/qux": detect.NormalFile,
			})

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

type mockResolver map[string]string

func (r mockResolver) List() []string {
	result := make([]string, len(r))
	i := 0
	for name := range r {
		result[i] = name
		i++
	}
	sort.Strings(result)
	return result
}

func (r mockResolver) Resolve(name string) (string, error) {
	return r[name], nil
}

type mockDetector map[string]detect.Type

func (d mockDetector) Detect(path string) detect.Type {
	return d[path]
}
