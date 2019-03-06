package warmup1

import (
	"strings"
	"testing"
)

func TestFront22Str(t *testing.T) {
	actual := front22("test")
	expected := "tetestte"

	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}

func TestFront22Str2(t *testing.T) {
	actual := front22("hi")
	expected := "hihihi"

	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}

func TestFront22Str1(t *testing.T) {
	actual := front22("x")
	expected := "xxx"

	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}
