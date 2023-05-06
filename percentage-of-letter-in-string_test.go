package leetcode

import "testing"

func testPercentageLetter(t *testing.T, s string, letter byte, expect int) {
	if actual := PercentageLetter(s, letter); actual != expect {
		t.Fatalf("expect=%d actual=%d", expect, actual)
	}
}

func TestPercentageLetter01(t *testing.T) {
	testPercentageLetter(t, "foobar", 'o', 33)
}

func TestPercentageLetter02(t *testing.T) {
	testPercentageLetter(t, "jjjj", 'k', 0)
}
