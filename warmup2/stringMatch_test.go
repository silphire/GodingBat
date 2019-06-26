package warmup2

import "testing"

func TestStringMatch1(t *testing.T) {
	expected := 3
	actual := stringMatch("xxcaazz", "xxbaaz")
	if expected != actual {
		t.Fatalf("expected %d, but actually %d", expected, actual)
	}
}

func TestStringMatch2(t *testing.T) {
	expected := 4
	actual := stringMatch("abcde", "abcdefg");
	if expected != actual {
		t.Fatalf("expected %d, but actually %d", expected, actual)
	}
}