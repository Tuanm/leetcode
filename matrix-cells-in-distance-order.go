package leetcode

import "sort"

// https://leetcode.com/problems/matrix-cells-in-distance-order/
func AllCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	result := make([][]int, 0, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result = append(result, []int{i, j})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return Abs(result[i][0]-rCenter)+Abs(result[i][1]-cCenter) < Abs(result[j][0]-rCenter)+Abs(result[j][1]-cCenter)
	})
	return result
}
