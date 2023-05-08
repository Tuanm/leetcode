package leetcode

// https://leetcode.com/problems/determine-if-two-events-have-conflict/
func HaveConflict(event1 []string, event2 []string) bool {
	return !((event1[0] < event2[0] && event1[1] < event2[1] && event1[1] < event2[0]) ||
		(event1[0] > event2[0] && event1[1] > event2[1] && event1[0] > event2[1]))
}
