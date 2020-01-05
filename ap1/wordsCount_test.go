package ap1

import "testing"

func TestWordsCount(t *testing.T) {
	if wordsCount([]string{"a", "bb", "b", "ccc"}, 1) != 2 {
		t.Fatal("expected 2 but actual is not")
	}
	if wordsCount([]string{"a", "bb", "b", "ccc"}, 3) != 1 {
		t.Fatal("expected 1 but actual is not")
	}
	if wordsCount([]string{"a", "bb", "b", "ccc"}, 4) != 0 {
		t.Fatal("expected 0 but actual is not")
	}
}
