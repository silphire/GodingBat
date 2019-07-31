package string1

import "testing"
import "strings"

func TestMiddleTwo4(t *testing.T) {
	expected := "ea"
	actual := middleTwo("meat")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestMiddleTwo6(t *testing.T) {
	expected := "dd"
	actual := middleTwo("middle")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}