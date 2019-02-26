package warmup1

import "testing"

func TestFirstIsOnlyTeen(t *testing.T) {
	actual := loneTeen(14, 21)
	if !actual {
		t.Fatal("expected true but actually not")
	}
}
func TestSecondIsOnlyTeen(t *testing.T) {
	actual := loneTeen(84, 16)
	if !actual {
		t.Fatal("expected true but actually not")
	}
}
func TestBothAreTeens(t *testing.T) {
	actual := loneTeen(19, 16)
	if actual {
		t.Fatal("expected false but actually not")
	}
}
func TestNooneIsTeen(t *testing.T) {
	actual := loneTeen(11, 22)
	if actual {
		t.Fatal("expected false but actually not")
	}
}
