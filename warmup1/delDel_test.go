package warmup1

import (
	"strings"
	"testing"
)

func TestDelDel(t *testing.T) {
	expected := "abc"
	actual := delDel("adelbc")

	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}

func TestDelLatterDel(t *testing.T) {
	expected := "abdelc"
	actual := delDel("abdelc")

	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
