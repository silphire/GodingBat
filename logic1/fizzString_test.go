package logic1

import "testing"
import "strings"

func TestFizzString(t *testing.T) {
	if strings.Compare(fizzString("fig"), "Fizz") != 0 {
		t.Fatal("expected Fizz but actually not")
	}

	if strings.Compare(fizzString("dib"), "Buzz") != 0 {
		t.Fatal("expected Buzz but actually not")
	}

	if strings.Compare(fizzString("fib"), "FizzBuzz") != 0 {
		t.Fatal("expected Buzz but actually not")
	}
}
