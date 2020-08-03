package repeatCoding

import (
	"fmt"
	"testing"
)

func TestInsterionSort(t *testing.T) {
	var arr = []int{2, 4, 1, 5, 7}
	insertionSort(arr)
	fmt.Printf("插入排序结果为：%v\n", arr)
}

//插入排序
func insertionSort(arr []int) {
	var size = len(arr)
	for i := 0; i < size-1; i++ {
		j := i + 1
		for j >= 1 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}
