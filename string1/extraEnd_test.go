package string1

import "testing"
import "strings"

func TestExtraEnd1(t *testing.T) {
	expected := "lololo"
	actual := extraEnd("Hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestExtraEnd2(t *testing.T) {
	expected := "AbAbAb"
	actual := extraEnd("Ab")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}