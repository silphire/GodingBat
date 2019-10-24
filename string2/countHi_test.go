package string2

import "testing"

func TestCountHi(t *testing.T) {
	if countHi("abc hi ho") != 1 {
		t.Fatal("expected 1 but actually not")
	}
	if countHi("ABChi hi") != 2 {
		t.Fatal("expected 1 but actually not")
	}
	if countHi("hihi") != 2 {
		t.Fatal("expected 1 but actually not")
	}
}
