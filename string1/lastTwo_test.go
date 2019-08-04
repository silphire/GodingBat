package string1

import "testing"
import "strings"

func TestLastTwo1(t *testing.T) {
	expected := "helol"
	actual := lastTwo("hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}