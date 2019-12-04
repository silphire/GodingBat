package warmup1

import (
	"strings"
	"testing"
)

func TestStartOz(t *testing.T) {
	expected := "oz"
	actual := startOz("ozymandias")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "z"
	actual = startOz("bzoo")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "o"
	actual = startOz("oxx")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
