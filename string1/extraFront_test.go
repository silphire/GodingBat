package string1

import "testing"
import "strings"

func TestExtraFront1(t *testing.T) {
	expected := "HeHeHe"
	actual := extraFront("Hello")

	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}