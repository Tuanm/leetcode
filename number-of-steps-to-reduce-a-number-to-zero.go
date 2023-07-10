package leetcode

func NumberOfSteps(num int) int {
	var steps int
	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		steps++
	}
	return steps
}
