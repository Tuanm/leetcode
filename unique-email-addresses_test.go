package leetcode

import "testing"

func testNumUniqueEmails(t *testing.T, input []string, expect int) {
	if actual := NumUniqueEmails(input); actual != expect {
		t.Fatalf("expect=%d actual=%d", expect, actual)
	}
}

func TestNumUniqueEmails01(t *testing.T) {
	testNumUniqueEmails(t, []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}, 2)
}

func TestNumUniqueEmails02(t *testing.T) {
	testNumUniqueEmails(t, []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}, 3)
}
