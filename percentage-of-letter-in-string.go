package leetcode

import "math"

func PercentageLetter(s string, letter byte) int {
	var c int = 0
	for _, ch := range s {
		if byte(ch) == letter {
			c++
		}
	}
	return int(math.Floor(float64(c*100) / float64(len(s))))
}
