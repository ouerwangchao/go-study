package array

import (
	"testing"
)

func TestMostWater(t *testing.T) {
	var arr = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	size := len(arr)
	area := 0
	i, j := 0, size-1
	for i < j {
		area = getMax(area, (j-i)*getMin(arr[i], arr[j]))
		if arr[i] > arr[j] {
			j--
		} else {
			i++
		}
	}
	print(area)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
