package status_test

import (
	"fmt"

	"github.com/ooesili/qfi/detect"
	. "github.com/ooesili/qfi/status"

	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Status", func() {
	Context("when given a single target", func() {
		It("displays it's status", func() {
			// setup test cases
			tests := []struct {
				fileType   detect.Type
				typeString string
				arrowChar  rune
			}{
				{detect.NormalFile, "NormalFile", '-'},
				{detect.UnwritableFile, "UnwritableFile", '#'},
				{detect.InaccessibleFile, "InaccessibleFile", '#'},
				{detect.NonexistentFile, "NonexistentFile", '-'},
				{detect.NormalDirectory, "NormalDirectory", '/'},
				{detect.UnreadableDirectory, "UnreadableDirectory", '/'},
				{detect.UnknownFile, "UnknownFile", '?'},
			}

			// run tests
			for _, test := range tests {
				resolver := mockResolver(map[string]string{
					"foobar": "/foo/bar",
				})
				detector := mockDetector(map[string]detect.Type{
					"/foo/bar": test.fileType,
				})
				status := Status(resolver, detector)
				Expect(status).To(Equal(
					fmt.Sprintf("foobar %c> /foo/bar\n", test.arrowChar)),
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

			status := Status(resolver, detector)
			Expect(status).To(Equal(`
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
