package string1

import "testing"
import "strings"

func TestFirstTwo1(t *testing.T) {
	expected := "He"
	actual := firstTwo("Hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestFirstTwo2(t *testing.T) {
	expected := "X"
	actual := firstTwo("X")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestFirstTwo3(t *testing.T) {
	expected := ""
	actual := firstTwo("")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}