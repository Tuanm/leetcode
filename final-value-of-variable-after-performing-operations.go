package leetcode

// https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
func FinalValueAfterOperations(operations []string) int {
	value := 0
	for _, operation := range operations {
		if operation == "X++" || operation == "++X" {
			value++
		} else {
			value--
		}
	}
	return value
}
