package string2

import (
	"strings"
	"testing"
)

func TestRepeatFront(t *testing.T) {
	expected := "ChocChoChC"
	actual := repeatFront("Chocolate", 4)
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
	
	expected = "ChoChC"
	actual = repeatFront("Chocolate", 3)
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
	
	expected = "IcI"
	actual = repeatFront("Ice Cream", 2)
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
