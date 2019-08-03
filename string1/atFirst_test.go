package string1

import "strings"
import "testing"

func TestAtFirst(t *testing.T) {
	expected := "he"
	actual := atFirst("hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestAtFirst2(t *testing.T) {
	expected := "x@"
	actual := atFirst("x")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestAtFirst3(t *testing.T) {
	expected := "@@"
	actual := atFirst("")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}