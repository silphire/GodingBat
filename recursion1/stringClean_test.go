package recursion1

import (
	"strings"
	"testing"
)

func TestStringClean(t *testing.T) {
	expected := "yza"
	actual := stringClean("yyzzza")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "abcd"
	actual = stringClean("abbbcdd")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "Helo"
	actual = stringClean("Hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
