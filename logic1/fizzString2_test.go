package logic1

import (
	"strings"
	"testing"
)

func TestFizzString2(t *testing.T) {
	if strings.Compare(fizzString2(1), "1!") != 0 {
		t.Fatal("expected \"1!\" but actually not")
	}
	if strings.Compare(fizzString2(2), "2!") != 0 {
		t.Fatal("expected \"2!\" but actually not")
	}
	if strings.Compare(fizzString2(3), "Fizz!") != 0 {
		t.Fatal("expected \"Fizz!\" but actually not")
	}
}
