package warmup1

import "testing"

func TestFirstIsTeen(t *testing.T) {
	if !hasTeen(13, 10, 20) {
		t.Fatal("first parameter has teen but testee says not so")
	}
}
func TestSecondIsTeen(t *testing.T) {
	if !hasTeen(21, 19, 113) {
		t.Fatal("second parameter has teen but testee says not so")
	}
}
func TestThirdIsTeen(t *testing.T) {
	if !hasTeen(21, 29, 14) {
		t.Fatal("second parameter has teen but testee says not so")
	}
}
func TestNoOneIsTeen(t *testing.T) {
	if hasTeen(21, 1, 113) {
		t.Fatal("no parameter has teen but testee says not so")
	}
}
