package array

import (
	"testing"
)

var arr = []int{1, 4, 2, 6, 3, 1, 4, 8}

func TestArrayDeduplication(t *testing.T) {
	mapDeduplication()
	t.Log(arr)
}

/**
map集合处理
*/
func mapDeduplication() {
	var m = make(map[int]int)
	var res = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		_, exist := m[arr[i]]
		if !exist {
			res = append(res, arr[i])
		}
		m[arr[i]] = arr[i]
	}
	arr = res
	return
}
