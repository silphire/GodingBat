package string2

import (
	"strings"
	"testing"
)

func TestZipZap(t *testing.T) {
	expected := "zpXzp"
	actual := zipZap("zipXzap")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "zpzp"
	actual = zipZap("zopzop")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	expected = "zzzpzp"
	actual = zipZap("zzzopzop")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
