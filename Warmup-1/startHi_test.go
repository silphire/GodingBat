package main

import "testing"

func TestBeginWithHi(t *testing.T) {
	if !startHi("hi there") {
		t.Fatal("expected true, but actually false")
	}
}
func TestJustHi(t *testing.T) {
	if !startHi("hi") {
		t.Fatal("expected true, but actually false")
	}
}

func TestNotHi(t *testing.T) {
	if startHi("hello hi") {
		t.Fatal("expected false, but actually true")
	}
}
