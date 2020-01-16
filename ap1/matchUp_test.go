package ap1

import "testing"

func TestMatchUp(t *testing.T) {
	actual := matchUp([]string{"aa", "bb", "cc"}, []string{"aaa", "xx", "bb"})
	if actual != 1 {
		t.Fatalf("expected 1 but actual is %d", actual)
	}

	actual = matchUp([]string{"aa", "bb", "cc"}, []string{"aaa", "b", "bb"})
	if actual != 2 {
		t.Fatalf("expected 2 but actual is %d", actual)
	}

	actual = matchUp([]string{"aa", "bb", "cc"}, []string{"", "", "ccc"})
	if actual != 1 {
		t.Fatalf("expected 1 but actual is %d", actual)
	}
}
