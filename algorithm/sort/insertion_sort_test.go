package sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	var arr = []int{1, 3, 2, 4, 2, 6, 8, 4}
	insertionSort(arr)
	fmt.Printf("插入排序结果为：%v\n", arr)
}

// insertionSort 当前值与已排序的值轮流比较
func insertionSort(arr []int) {
	var length = len(arr)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
}
