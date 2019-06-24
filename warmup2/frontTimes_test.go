package warmup2

import (
	"strings"
	"testing"
)

func TestFrontTimes(t *testing.T) {
	expected := "foofoo"
	actual := frontTimes("foobarbaz", 2)
	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}

func TestFrontTimes2(t *testing.T) {
	expected := "foofoofoo"
	actual := frontTimes("foobarbaz", 3)
	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}
