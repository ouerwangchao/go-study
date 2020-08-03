package repeatCoding

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var arr = []int{3, 0, 1, 4, 6, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Printf("快速排序结果为：%v\n", arr)
	fmt.Print("--------------------\n")
	var arrs = []int{3, 0, 1, 4, 6, 2}
	quickSorts(arrs, 0, len(arrs)-1)
	fmt.Printf("快速排序结果为：%v\n", arr)
}

func quickSorts(arr []int, left, right int) {
	if left > right {
		return
	}
	var l, r = -1 + left, left //左右指针
	var mid = arr[right]
	for i := r; i < right; i++ {
		if arr[i] < mid {
			l++
			arr[l], arr[i] = arr[i], arr[l]
		}
	}
	arr[l+1], arr[right] = arr[right], arr[l+1]
	if l > left {
		quickSorts(arr, left, l)
		fmt.Printf("sorts-left: left is %v, right is %v,arr is %v\n", left, l, arr)
	}
	if right > l+2 {
		quickSorts(arr, l+2, right)
		fmt.Printf("sorts-right: left is %v, right is %v,arr is %v\n", l+2, right, arr)
	}
}

func quickSort(arr []int, left, right int) {
	if left > right {
		return
	}
	var l, r = -1, left //左右指针
	var mid = arr[right]
	for i := r; i < right; i++ {
		if arr[i] < mid {
			l++
			arr[l], arr[i] = arr[i], arr[l]
		}
	}
	arr[l+1], arr[right] = arr[right], arr[l+1]
	if l > 0 {
		quickSort(arr[:l+1], 0, l)
		fmt.Printf("sort-left: left is %v, right is %v, arr is %v \n", 0, l, arr[:l+1])
	}
	if right-l-2 > 0 {
		quickSort(arr[l+2:], 0, right-l-2)
		fmt.Printf("sort-right: left is %v, right is %v, arr is %v \n", 0, right-l-2, arr[l+2:])
	}
}
