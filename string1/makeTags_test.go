package string1

import "testing"
import "strings"

func TestMakeTags(t *testing.T) {
	expected := "<s>hello</s>"
	actual := makeTags("s", "hello")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}