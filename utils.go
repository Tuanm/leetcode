package leetcode

func Abs[T int | float64](i T) T {
	if i >= 0 {
		return i
	}
	return -i
}
