package warmup1

import "testing"

func TestBoth30s(t *testing.T) {
	actual := in3050(30, 39)
	if !actual {
		t.Fatal("expected true but actually not")
	}
}

func TestOtherGroup(t *testing.T) {
	actual := in3050(35, 45)
	if actual {
		t.Fatal("expected false but actually not")
	}
}

func TestBoth40s(t *testing.T) {
	actual := in3050(49, 40)
	if !actual {
		t.Fatal("expected true but actually not")
	}
}
