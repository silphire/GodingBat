package string3

import (
	"strings"
	"testing"
)

func TestNotReplace(t *testing.T) {
	expected := "is not test"
	actual := notReplace("is test")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "is not-is not"
	actual = notReplace("is-is")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "This is not right"
	actual = notReplace("This is right")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
