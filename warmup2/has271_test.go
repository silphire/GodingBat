package warmup2

import "testing"

func TestContains271(t *testing.T) {
	expected := true
	actual := has271([]int{1, 2, 7, 1})
	if expected != actual {
		t.Fatalf("expected true, but actual is false")
	}
}

func TestNo271(t *testing.T) {
	expected := false
	actual := has271([]int{1, 2, 8, 1})
	if expected != actual {
		t.Fatalf("expected false, but actual is true")
	}
}
