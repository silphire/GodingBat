package array2

import (
	"strings"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	expected := []string{"1", "2", "Fizz", "4", "Buzz"}
	actual := fizzBuzz(1, 6)
	if len(expected) != len(actual) {
		t.Fatalf("size mismatch, expected:%d actual:%d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	expected = []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7"}
	actual = fizzBuzz(1, 8)
	if len(expected) != len(actual) {
		t.Fatalf("size mismatch, expected:%d actual:%d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	expected = []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"}
	actual = fizzBuzz(1, 11)
	if len(expected) != len(actual) {
		t.Fatalf("size mismatch, expected:%d actual:%d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}
}
