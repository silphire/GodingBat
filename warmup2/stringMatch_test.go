package warmup2

import "testing"

func TestStringMatch1(t *testing.T) {
	expected := 3
	actual := stringMatch("xxcaazz", "xxbaaz")
	if expected != actual {
		t.Fatal("expected %d, but actually %d", expected, actual)
	}
}
