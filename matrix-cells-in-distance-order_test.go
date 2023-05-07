package leetcode

import (
	"reflect"
	"testing"
)

func testAllCellsDistOrder(t *testing.T, rows int, cols int, rCenter int, cCenter int, expect [][]int) {
	if actual := AllCellsDistOrder(rows, cols, rCenter, cCenter); !reflect.DeepEqual(actual, expect) {
		t.Fatalf("expect=%d actual=%d", expect, actual)
	}
}

func TestAllCellsDistOrder01(t *testing.T) {
	testAllCellsDistOrder(t, 1, 2, 0, 0, [][]int{{0, 0}, {0, 1}})
}

func TestAllCellsDistOrder02(t *testing.T) {
	testAllCellsDistOrder(t, 2, 2, 0, 1, [][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}})
}

func TestAllCellsDistOrder03(t *testing.T) {
	testAllCellsDistOrder(t, 2, 3, 1, 2, [][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}})
}
