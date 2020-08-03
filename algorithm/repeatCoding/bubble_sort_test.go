package repeatCoding

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var arr = []int{1, 2, 3, 1, 4, 5, 6}
	bubbleSort(arr)
	fmt.Printf("冒泡排序结果为：%v\n", arr)
}

//冒泡排序
func bubbleSort(arr []int) {
	var size = len(arr)
	for i := 0; i < size; i++ {
		var t = true
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				t = false
			}
		}
		fmt.Printf("i=%v\n", i)
		if t {
			break
		}
	}
}
