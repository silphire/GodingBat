package main

import "testing"

func testEqual(t *testing.T) {
	if !nearHundred(100) {
		t.Fatal("expected true, but got false")
	}
	if !nearHundred(200) {
		t.Fatal("expected true, but got false")
	}
}

func testWithin(t *testing.T) {
	if !nearHundred(105) {
		t.Fatal("expected true, but got false")
	}
	if !nearHundred(98) {
		t.Fatal("expected true, but got false")
	}
	if !nearHundred(203) {
		t.Fatal("expected true, but got false")
	}
	if !nearHundred(194) {
		t.Fatal("expected true, but got false")
	}
}

func testUnder(t *testing.T) {
	if nearHundred(89) {
		t.Fatal("exected false, but got true")
	}
	if nearHundred(154) {
		t.Fatal("exected false, but got true")
	}
}

func testOver(t *testing.T) {
	if nearHundred(111) {
		t.Fatal("exected false, but got true")
	}
	if nearHundred(222) {
		t.Fatal("exected false, but got true")
	}
}
