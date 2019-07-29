package string1

import "testing"
import "strings"

func TestComboString1(t *testing.T) {
	expected := "hiHellohi"
	actual := comboString("Hello", "hi")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}