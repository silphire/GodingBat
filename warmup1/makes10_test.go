package warmup1

import "testing"

func TestOneIs10(t *testing.T) {
	actual := makes10(5, 10)
	if !actual {
		t.Fatal("expected true but actually not")
	}
}

func TestNeither10(t *testing.T) {
	actual := makes10(5, 8)
	if actual {
		t.Fatal("expected false but actually not")
	}
}

func TestArgsMake10(t *testing.T) {
	actual := makes10(2, 8)
	if !actual {
		t.Fatal("expected true but actually not")
	}
}
