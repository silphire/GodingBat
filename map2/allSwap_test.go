package map2

import (
	"reflect"
	"testing"
)

func TestAllSwap(t *testing.T) {
	actual := allSwap([]string{"ab", "ac"})
	expected := []string{"ac", "ab"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = allSwap([]string{"ax", "bx", "cx", "cy", "by", "ay", "aaa", "azz"})
	expected = []string{"ay", "by", "cy", "cx", "bx", "ax", "azz", "aaa"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}

	actual = allSwap([]string{"ax", "bx", "ay", "by", "ai", "aj", "bx", "by"})
	expected = []string{"ay", "by", "ax", "bx", "aj", "ai", "by", "bx"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("not equal")
	}
}
