package warmup1

import "testing"

func TestOneE(t *testing.T) {
	if !stringE("Hello") {
		t.Fatal("expected true but actually not")
	}
}
func TestThreeEs(t *testing.T) {
	if !stringE("Hellee") {
		t.Fatal("expected true but actually not")
	}
}
func TestFourEs(t *testing.T) {
	if stringE("Helelee") {
		t.Fatal("expected false but actually not")
	}
}
func TestNoEs(t *testing.T) {
	if stringE("Hllo") {
		t.Fatal("expected false but actually not")
	}
}
