package exhaustive

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestIsGeneratedFile(t *testing.T) {
	t.Run("generated", func(t *testing.T) {
		const source = `
	// Code generated by some thing. DO NOT EDIT.

package foo
func f() {
}`
		f, err := parser.ParseFile(token.NewFileSet(), "", source, parser.ParseComments)
		assertNoError(t, err)
		if !isGeneratedFile(f) {
			t.Errorf("unexpectedly false")
		}
	})

	t.Run("not generated", func(t *testing.T) {
		const source = `package foo
func f() {
}`
		f, err := parser.ParseFile(token.NewFileSet(), "", source, parser.ParseComments)
		assertNoError(t, err)
		if isGeneratedFile(f) {
			t.Errorf("unexpectedly true")
		}
	})
}
