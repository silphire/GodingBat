package warmup1

import (
	"strings"
	"testing"
)

func TestFront3String(t *testing.T) {
	actual := front3("Golang")
	expected := "GolGolGol"
	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}

func TestFront3String3(t *testing.T) {
	actual := front3("xyz")
	expected := "xyzxyzxyz"
	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}
