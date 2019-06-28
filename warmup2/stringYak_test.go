package warmup2

import (
	"strings"
	"testing"
)

func TestYakFound(t *testing.T) {
	expected := "foo"
	actual := stringYak("yakfoo")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}
