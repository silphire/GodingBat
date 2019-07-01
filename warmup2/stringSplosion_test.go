package warmup2

import (
	"strings"
	"testing"
)

func TestStringSplosionAbc(t *testing.T) {
	expected := "aababc"
	actual := stringSplosion("abc")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}
