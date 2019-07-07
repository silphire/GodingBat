package string1

import (
	"strings"
	"testing"
)

func TestNonStart1(t *testing.T) {
	expected := "ellohere"
	actual := nonStart("Hello", "There")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}
