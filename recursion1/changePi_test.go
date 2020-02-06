package recursion1

import (
	"strings"
	"testing"
)

func TestChangePi(t *testing.T) {
	if strings.Compare(changePi("xpix"), "x3.14x") != 0 {
		t.Fatal("expected \"x3.14x\" but actual is not")
	}
	if strings.Compare(changePi("pipi"), "3.143.14") != 0 {
		t.Fatal("expected \"3.143.14\" but actual is not")
	}
	if strings.Compare(changePi("pip"), "3.14p") != 0 {
		t.Fatal("expected \"3.14p\" but actual is not")
	}
}
