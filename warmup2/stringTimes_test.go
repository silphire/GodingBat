package warmup2

import (
	"strings"
	"testing"
)

func TestStringTimes(t *testing.T) {
	actual := stringTimes("hi", 3)
	expected := "hihihi"
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}
