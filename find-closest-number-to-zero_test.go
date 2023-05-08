package leetcode

import "testing"

func testFindClosestNumber(t *testing.T, nums []int, expect int) {
	if actual := FindClosestNumber(nums); actual != expect {
		t.Fatalf("expect=%d actual=%d", expect, actual)
	}
}

func TestFindClosestNumber01(t *testing.T) {
	testFindClosestNumber(t, []int{-4, -2, 1, 4, 8}, 1)
}

func TestFindClosestNumber02(t *testing.T) {
	testFindClosestNumber(t, []int{2, -1, 1}, 1)
}

func TestFindClosestNumber03(t *testing.T) {
	testFindClosestNumber(t, []int{0}, 0)
}
