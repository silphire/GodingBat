package warmup1

import "testing"

func TestTemp1IsHot(t *testing.T) {
	actual := icyHot(110, -100)
	if !actual {
		t.Fatal("expected true but actually not")
	}
}
func TestTemp2IsHot(t *testing.T) {
	actual := icyHot(-110, 120)
	if !actual {
		t.Fatal("expected true but actually not")
	}
}
func TestBothAreHot(t *testing.T) {
	actual := icyHot(10, 120)
	if actual {
		t.Fatal("expected false but actually not")
	}
}
func TestBothAreCold(t *testing.T) {
	actual := icyHot(-10, -5)
	if actual {
		t.Fatal("expected false but actually not")
	}
}
