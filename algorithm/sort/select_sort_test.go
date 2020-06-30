package sort

import "testing"

func TestSelectSort(t *testing.T) {
	var arr = []int{1, 3, 1, 4, 6, 2, 7, 3}
	res := SelectSort(arr)
	t.Log(res)
}

func SelectSort(array []int) []int {
	size := len(array)
	for i := 0; i < size; i++ { //每次默认array[i]为最小值，与后续值一一比较
		for j := i + 1; j < size; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}
