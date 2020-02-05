package recursion1

import (
	"strings"
	"testing"
)

func TestChangeXY(t *testing.T) {
	if strings.Compare(changeXY("codex"), "codey") != 0 {
		t.Fatal("expected \"codey\" but actual is not")
	}
	if strings.Compare(changeXY("xxhixx"), "yyhiyy") != 0 {
		t.Fatal("expected \"yyhiyy\" but actual is not")
	}
	if strings.Compare(changeXY("xhixhix"), "yhiyhiy") != 0 {
		t.Fatal("expected \"yhiyhiy\" but actual is not")
	}
}
