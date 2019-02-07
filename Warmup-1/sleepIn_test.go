package main

import "testing"

func TestWeekdayVacation(t *testing.T) {
	result := sleepIn(true, true)
	if !result {
		t.Fatal("bad result")
	}
}

func TestHolidayNotVacation(t *testing.T) {
	result := sleepIn(true, true)
	if !result {
		t.Fatal("bad result")
	}
}

func TestHolidayVacation(t *testing.T) {
	result := sleepIn(true, true)
	if !result {
		t.Fatal("bad result")
	}
}

func TestWeekdayNotVacation(t *testing.T) {
	result := sleepIn(true, false)
	if result {
		t.Fatal("bad result")
	}
}
