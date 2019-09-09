package logic1

import "testing"

func TestLove6(t *testing.T) {
	if !love6(6, 4) {
		t.Fatal("expected true but actually noy")
	}
	if love6(5, 4) {
		t.Fatal("expected false but actually noy")
	}
	if !love6(5, 1) {
		t.Fatal("expected true but actually noy")
	}
	if !love6(2, 8) {
		t.Fatal("expected true but actually noy")
	}
}