package recursion1

import "testing"

func TestCountPairs(t *testing.T) {
	if countPairs("axa") != 1 {
		t.Fatal("expected 1 but actual is not")
	}

	if countPairs("axax") != 2 {
		t.Fatal("expected 2 but actual is not")
	}
	
	if countPairs("axbx") != 1 {
		t.Fatal("expected 1 but actual is not")
	}
}
