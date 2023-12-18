package leetcode

import "testing"

func testSmallestChair(t *testing.T, times [][]int, targetFriend int, expect int) {
	if actual := SmallestChair(times, targetFriend); actual != expect {
		t.Fatalf("expect=%d actual=%d", expect, actual)
	}
}

func TestSmallestChair01(t *testing.T) {
	testSmallestChair(t, [][]int{
		{1, 4},
		{2, 3},
		{4, 6},
	}, 1, 1)
}

func TestSmallestChair02(t *testing.T) {
	testSmallestChair(t, [][]int{
		{3, 10},
		{1, 5},
		{2, 6},
	}, 0, 2)
}
