package warmup2

import (
	"strings"
	"testing"
)

func TestAltPairs(t *testing.T) {
	actual := altPairs("kitten")
	expected := "kien"
	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}

func TestAltPairs2(t *testing.T) {
	actual := altPairs("Chocolate")
	expected := "Chole"
	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}

func TestAltPairs3(t *testing.T) {
	actual := altPairs("CodingHorror")
	expected := "Congrr"
	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected \"%s\", but actually \"%s\"", expected, actual)
	}
}
