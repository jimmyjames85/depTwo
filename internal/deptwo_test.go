package internal

import (
	"testing"
)

func TestOne(t *testing.T) {
	expect := "This is Two: One is: One"
	actual := Two()
	if expect != actual {
		t.Fatalf("expect: %q actual: %q", expect, actual)
	}
}
