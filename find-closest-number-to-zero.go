package leetcode

// https://leetcode.com/problems/find-closest-number-to-zero/
func FindClosestNumber(nums []int) int {
	var closest int = nums[0]
	for _, num := range nums {
		diff := Abs(num) - Abs(closest)
		if (diff < 0) || (diff == 0 && num > closest) {
			closest = num
		}
	}
	return closest
}
