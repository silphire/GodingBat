package logic1

import (
	"strings"
	"testing"
)

func TestAlarmClock(t *testing.T) {
	if strings.Compare(alarmClock(1, false), "7:00") != 0 {
		t.Fatal("expected \"07:00\" but actually not")
	}
	if strings.Compare(alarmClock(5, false), "7:00") != 0 {
		t.Fatal("expected \"07:00\" but actually not")
	}
	if strings.Compare(alarmClock(0, false), "10:00") != 0 {
		t.Fatal("expected \"10:00\" but actually not")
	}
}
