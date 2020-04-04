package map2

import (
	"strings"
	"testing"
)

func TestFirstSwap(t *testing.T) {
	actual := firstSwap([]string{"ab", "ac"})
	expected := []string{"ac", "ab"}
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	actual = firstSwap([]string{"ax", "bx", "cx", "cy", "by", "ay", "aaa", "azz"})
	expected = []string{"ay", "by", "cy", "cx", "bx", "ax", "aaa", "azz"}
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	actual = firstSwap([]string{"ax", "bx", "ay", "by", "ai", "aj", "bx", "by"})
	expected = []string{"ay", "by", "ax", "bx", "ai", "aj", "bx", "by"}
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}
}
