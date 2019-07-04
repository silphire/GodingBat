package string1

import "testing"
import "strings"

func TestHelloName1(t *testing.T) {
	expected := "Hello Bob!"
	actual := helloName("Bob")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestHelloName2(t *testing.T) {
	expected := "Hello Alice!"
	actual := helloName("Alice")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

