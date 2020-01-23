package ap1

import "testing"

func TestUserCompare(t *testing.T) {
	if (userCompare("bb", 1, "zz", 2)) != -1 {
		t.Fatal("expected -1 but actual is not")
	}
	if (userCompare("bb", 1, "aa", 2)) != 1 {
		t.Fatal("expected 1 but actual is not")
	}
	if (userCompare("bb", 1, "bb", 1)) != 0 {
		t.Fatal("expected 0 but actual is not")
	}
}
