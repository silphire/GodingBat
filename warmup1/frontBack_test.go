package warmup1

import (
	"strings"
	"testing"
)

func TestSwap(t *testing.T) {
	actual := frontBack("back")
	expected := "kacb"
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestOneChar(t *testing.T) {
	expected := frontBack("x")
	actual := "x"
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestOddChars(t *testing.T) {
	actual := frontBack("abcde")
	expected := "ebcda"
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
