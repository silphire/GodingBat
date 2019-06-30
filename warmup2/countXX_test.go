package warmup2

import "testing"

func TestCountOneXX(t *testing.T) {
	if countXX("axxb") != 1 {
		t.Fatal("expected 1 but actually not")
	}
}

func TestCountTwoXX(t *testing.T) {
	if countXX("xxx") != 2 {
		t.Fatal("expected 1 but actually not")
	}
}

func TestCountThreeXX(t *testing.T) {
	if countXX("xxxx") != 3 {
		t.Fatal("expected 1 but actually not")
	}
}

func TestCountNoXX(t *testing.T) {
	if countXX("axbx") != 0 {
		t.Fatal("expected 1 but actually not")
	}
}
