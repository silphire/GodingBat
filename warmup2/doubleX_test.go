package warmup2

import "testing"

func TestFoundDoubleX(t *testing.T) {
	if !doubleX("aaxxbb") {
		t.Fatal("expected true but not")
	}
}

func TestNotFoundDoubleX(t *testing.T) {
	if doubleX("axbxcx") {
		t.Fatal("expected false but not")
	}
}
