package leetcode

import "testing"

func testNumberOfSteps(t *testing.T, num int, expect int) {
	if actual := NumberOfSteps(num); actual != expect {
		t.Fatalf("expect=%d actual=%d", expect, actual)
	}
}

func TestNumberOfSteps01(t *testing.T) {
	testNumberOfSteps(t, 14, 6)
}

func TestNumberOfSteps02(t *testing.T) {
	testNumberOfSteps(t, 8, 4)
}

func TestNumberOfSteps03(t *testing.T) {
	testNumberOfSteps(t, 123, 12)
}
