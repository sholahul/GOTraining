package hash

import "testing"

func TestSHA256(t *testing.T) {
	pwd := NewSHA256("Test")
	t.Log(pwd)
}
