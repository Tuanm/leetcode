package leetcode

import "testing"

func testHaveConflict(t *testing.T, event1 []string, event2 []string, expect bool) {
	if actual := HaveConflict(event1, event2); actual != expect {
		t.Fatalf("expect=%t actual=%t", expect, actual)
	}
}

func TestHaveConflict01(t *testing.T) {
	testHaveConflict(t, []string{"01:15", "02:00"}, []string{"02:00", "03:00"}, true)
}

func TestHaveConflict02(t *testing.T) {
	testHaveConflict(t, []string{"01:00", "02:00"}, []string{"01:20", "03:00"}, true)
}

func TestHaveConflict03(t *testing.T) {
	testHaveConflict(t, []string{"10:00", "11:00"}, []string{"14:00", "15:00"}, false)
}
