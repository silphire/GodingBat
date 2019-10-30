package string2

import (
	"strings"
	"testing"
)

func TestPlusOut(t *testing.T) {
	expected := "++xy++"
	actual := plusOut("12xy34", "xy")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "1+++++"
	actual = plusOut("12xy34", "1")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "++xy++xy+++xy"
	actual = plusOut("12xy34xyabcxy", "xy")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
