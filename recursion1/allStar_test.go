package recursion1

import (
	"strings"
	"testing"
)

func TestAllStar(t *testing.T) {
	expected := "h*e*l*l*o"
	actual := allStar("hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "a*b*c"
	actual = allStar("abc")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "a*b"
	actual = allStar("ab")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
