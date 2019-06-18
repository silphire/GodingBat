package warmup2

import "testing"

func TestFoundDoubleX(t *testing.T) {
	if !doubleX("aaxxbb") {
		t.Fatal("expected true but not")
	}
}
