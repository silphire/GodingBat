package warmup1

import (
	"strings"
	"testing"
)

func TestNormalWord(t *testing.T) {
	actual := notString("good")
	expected := "not good"

	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}

func TestWordPrefixByNot(t *testing.T) {
	actual := notString("not good")
	expected := "not good"

	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}
