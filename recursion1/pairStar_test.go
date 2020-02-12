package recursion1

import (
	"strings"
	"testing"
)

func TestPairStar(t *testing.T) {
	expected := "hel*lo"
	actual := pairStar("hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "x*xy*y"
	actual = pairStar("xxyy")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "a*a*a*a"
	actual = pairStar("aaaa")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
