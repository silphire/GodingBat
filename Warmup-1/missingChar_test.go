package main

import (
	"strings"
	"testing"
)

func TestMid(t *testing.T) {
	var result = missingChar("hello", 1)
	if strings.Compare(result, "hllo") != 0 {
		t.Fatalf("expected \"%s\", but \"%s\" actually", "hllo", result)
	}
}
func TestBegin(t *testing.T) {
	var result = missingChar("hello", 0)
	if strings.Compare(result, "ello") != 0 {
		t.Fatalf("expected \"%s\", but \"%s\" actually", "hllo", result)
	}
}
func TestEnd(t *testing.T) {
	var result = missingChar("hello", 4)
	if strings.Compare(result, "hell") != 0 {
		t.Fatalf("expected \"%s\", but \"%s\" actually", "hllo", result)
	}
}
