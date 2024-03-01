// Code generated by generator. DO NOT EDIT.

package t

import (
	"context"
	"testing"
)

// -----------------------------
// Free functions
// -----------------------------

func nonTestHelper(t int) {}

func helperWithoutHelper(t *testing.T) {} 

func helperWithHelper(t *testing.T) {
	t.Helper()
}

func helperWithEmptyStringBeforeHelper(t *testing.T) {

	t.Helper()
}

func helperWithHelperAfterAssignment(t *testing.T) { 
	_ = 0
	t.Helper()
}

func helperWithHelperAfterOtherCall(t *testing.T) { 
	ff()
	t.Helper()
}

func helperWithHelperAfterOtherSelectionCall(t *testing.T) { 
	t.Fail()
	t.Helper()
}

func helperParamNotFirst(s string, i int, t *testing.T) { // want "parameter \\*testing.T should be the first or after context.Context"
	t.Helper()
}

func helperParamSecondWithoutContext(s string, t *testing.T, i int) { // want "parameter \\*testing.T should be the first or after context.Context"
	t.Helper()
}

func helperParamSecondWithContext(ctx context.Context, t *testing.T) {
	t.Helper()
}

func helperWithIncorrectName(o *testing.T) { 
	o.Helper()
}

func helperWithAnonymousHelper(t *testing.T) {
	t.Helper()
	func(t *testing.T) {}(t) 
}

func helperWithNoName(_ *testing.T) {
}

// -----------------------------
// Methods of helper type
type helperType struct{}
// -----------------------------

func (h helperType) nonTestHelper(t int) {}

func (h helperType) helperWithoutHelper(t *testing.T) {} 

func (h helperType) helperWithHelper(t *testing.T) {
	t.Helper()
}

func (h helperType) helperWithEmptyStringBeforeHelper(t *testing.T) {

	t.Helper()
}

func (h helperType) helperWithHelperAfterAssignment(t *testing.T) { 
	_ = 0
	t.Helper()
}

func (h helperType) helperWithHelperAfterOtherCall(t *testing.T) { 
	ff()
	t.Helper()
}

func (h helperType) helperWithHelperAfterOtherSelectionCall(t *testing.T) { 
	t.Fail()
	t.Helper()
}

func (h helperType) helperParamNotFirst(s string, i int, t *testing.T) { // want "parameter \\*testing.T should be the first or after context.Context"
	t.Helper()
}

func (h helperType) helperParamSecondWithoutContext(s string, t *testing.T, i int) { // want "parameter \\*testing.T should be the first or after context.Context"
	t.Helper()
}

func (h helperType) helperParamSecondWithContext(ctx context.Context, t *testing.T) {
	t.Helper()
}

func (h helperType) helperWithIncorrectName(o *testing.T) { 
	o.Helper()
}

func (h helperType) helperWithAnonymousHelper(t *testing.T) {
	t.Helper()
	func(t *testing.T) {}(t) 
}

func (h helperType) helperWithNoName(_ *testing.T) {
}

func ff() {}
