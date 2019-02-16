package warmup1

import "testing"

func TestStartWithMix(t *testing.T) {
	if !mixStart("mix start") {
		t.Fatal("expected true, but false actually")
	}
}

func TestWithMixOnly(t *testing.T) {
	if !mixStart("mix start") {
		t.Fatal("expected true, but false actually")
	}
}

func TestStartWithPix(t *testing.T) {
	if !mixStart("pix start") {
		t.Fatal("expected true, but false actually")
	}
}
func TestStartWithPlz(t *testing.T) {
	if mixStart("plz start") {
		t.Fatal("expected false, but true actually")
	}
}
