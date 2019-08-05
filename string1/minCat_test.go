package string1

import "testing"
import "strings"

func TestMinCat1(t *testing.T) {
	expected := "loHi"
	actual := minCat("Hello", "Hi")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestMinCat2(t *testing.T) {
	expected := "Hilo"
	actual := minCat("Hi", "Hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestMinCat3(t *testing.T) {
	expected := "HiGo"
	actual := minCat("Hi", "Go")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}