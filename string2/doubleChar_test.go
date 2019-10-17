package string2

import (
	"strings"
	"testing"
)

func TestDoubleChar(t *testing.T) {
	if strings.Compare(doubleChar("The"), "TThhee") != 0 {
		t.Fatal("expected \"TThhee\" but actually not")
	}
	if strings.Compare(doubleChar("AAbb"), "AAAAbbbb") != 0 {
		t.Fatal("expected \"AAAAbbbb\" but actually not")
	}
	if strings.Compare(doubleChar("Hi-There"), "HHii--TThheerree") != 0 {
		t.Fatal("expected \"HHii--TThheerree\" but actually not")
	}
}
