package warmup1

import (
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	actual := backAround("Hello")
	expected := "oHelloo"
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}
