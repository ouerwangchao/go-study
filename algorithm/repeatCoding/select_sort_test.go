package repeatCoding

import (
	"fmt"
	"testing"
)

func TestSlectSort(t *testing.T) {
	var arr = []int{3, 1, 4, 6, 2, 6, 1, 5}
	selectSort(arr)
	fmt.Printf("选择排序结果为：%v\n", arr)
}

//选择排序
func selectSort(arr []int) {
	var size = len(arr)
	for i := 0; i < size-1; i++ {
		minIdx := i //默认第一个元素为最小值
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j //记录最小值的下标
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
