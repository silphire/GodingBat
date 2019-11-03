package string2

import (
	"strings"
	"testing"
)

func TestMixString(t *testing.T) {
	expected := "axbycz"
	actual := mixString("abc", "xyz")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "HTihere"
	actual = mixString("Hi", "There")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "xTxhxexre"
	actual = mixString("xxxx", "There")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
