package main

import "testing"

func TestOver21(t *testing.T) {
	result := diff21(23)
	if result != 2 {
		t.Fatal("bad result")
	}
}

func TestUnder21(t *testing.T) {
	result := diff21(13)
	if result != 8 {
		t.Fatal("bad result")
	}
}
