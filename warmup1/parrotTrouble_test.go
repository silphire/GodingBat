package warmup1

import "testing"

func TestTalkingEarly(t *testing.T) {
	result := parrotTrouble(true, 6)
	if !result {
		t.Fatal("expected true but false actually")
	}
}

func TestTalkingInTime(t *testing.T) {
	result := parrotTrouble(true, 7)
	if result {
		t.Fatal("expected false but true actually")
	}
}

func TestSilentInEarlyTime(t *testing.T) {
	result := parrotTrouble(false, 6)
	if result {
		t.Fatal("expected false but true actually")
	}
}

func TestTalkingNight(t *testing.T) {
	result := parrotTrouble(true, 21)
	if !result {
		t.Fatal("expected true but false actually")
	}
}

func TestSilentInNight(t *testing.T) {
	result := parrotTrouble(false, 21)
	if result {
		t.Fatal("expected false but true actually")
	}
}
