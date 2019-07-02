package warmup2

import (
	"strings"
	"testing"
)

func TestStringX(t *testing.T) {
	expected := "xHix"
	actual := stringX("xxHxix")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestStringX2(t *testing.T) {
	expected := "abcd"
	actual := stringX("abxxxcd")	
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}