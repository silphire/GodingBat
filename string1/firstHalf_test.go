package string1

import (
	"strings"
	"testing"
)

func TestFirstHalf1(t *testing.T) {
	expected := "foo"
	actual := firstHalf("foobar")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}
