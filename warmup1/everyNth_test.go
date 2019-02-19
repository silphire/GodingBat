package warmup1

import (
	"strings"
	"testing"
)

func TestSkipOneChar(t *testing.T) {
	result := everyNth("Hello", 2)
	if strings.Compare(result, "Hlo") != 0 {
		t.Fatalf("expected \"%s\" but \"%s\" actually", "Hlo", result)
	}
}

func TestEvenNumberLengthString(t *testing.T) {
	result := everyNth("good bye", 2)
	if strings.Compare(result, "go y") != 0 {
		t.Fatalf("expected \"%s\" but \"%s\" actually", "go y", result)
	}
}
func TestSkipThreeChars(t *testing.T) {
	result := everyNth("good bye", 3)
	if strings.Compare(result, "gdy") != 0 {
		t.Fatalf("expected \"%s\" but \"%s\" actually", "go y", result)
	}
}
