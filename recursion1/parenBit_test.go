package recursion1

import (
	"strings"
	"testing"
)

func TestParenBit(t *testing.T) {
	expected := "(abc)"
	actual := parenBit("xyz(abc)123")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "(hello)"
	actual = parenBit("x(hello)")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "(xy)"
	actual = parenBit("(xy)1")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
