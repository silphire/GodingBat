package recursion1

import "testing"

func TestCountAbc(t *testing.T) {
	if countAbc("abc") != 1 {
		t.Fatal("expected 1 but actual is not")
	}

	if countAbc("abcxxabc") != 2 {
		t.Fatal("expected 2 but actual is not")
	}

	if countAbc("abaxxaba") != 2 {
		t.Fatal("expected 2 but actual is not")
	}
}
