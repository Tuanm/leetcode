package leetcode

import "testing"

func testFinalValueAfterOperations(t *testing.T, input []string, expect int) {
	if actual := FinalValueAfterOperations(input); actual != expect {
		t.Fatalf("expect=%d actual=%d", expect, actual)
	}
}

func TestFinalValueAfterOperations01(t *testing.T) {
	testFinalValueAfterOperations(t, []string{"--X", "X++", "X++"}, 1)
}

func TestFinalValueAfterOperations02(t *testing.T) {
	testFinalValueAfterOperations(t, []string{"++X", "++X", "X++"}, 3)
}

func TestFinalValueAfterOperations03(t *testing.T) {
	testFinalValueAfterOperations(t, []string{"X++", "++X", "--X", "X--"}, 0)
}
