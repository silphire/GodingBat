package warmup1

import (
	"strings"
	"testing"
)

func TestEndUpLongString(t *testing.T) {
	actual := endUp("hello")
	expected := "heLLO"

	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}

func TestEndUpShortString(t *testing.T) {
	actual := endUp("hi")
	expected := "HI"

	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}
