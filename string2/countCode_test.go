package string2

import "testing"

func TestCountCode(t *testing.T) {
	actual := countCode("aaacodebbb")
	if actual != 1 {
		t.Fatalf("expected 1 but actual is %d", actual)
	}

	actual = countCode("codexxcode")
	if actual != 2 {
		t.Fatalf("expected 2 but actual is %d", actual)
	}
	
	actual = countCode("cozexxcope")
	if actual != 2 {
		t.Fatalf("expected 2 but actual is %d", actual)
	}
}
