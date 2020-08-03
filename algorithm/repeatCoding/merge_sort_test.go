package repeatCoding

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var nums1 = []int{2, 3, 5, 7}
	var nums2 = []int{3, 4, 6, 8}
	res := merge(nums1, nums2)
	fmt.Printf("v is %v\n", res)
}

func merge(nums1, nums2 []int) []int {
	var i, j = 0, 0
	var temp []int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			temp = append(temp, nums1[i])
			i++
		} else {
			temp = append(temp, nums2[j])
			j++
		}
	}
	for i < len(nums1) {
		temp = append(temp, nums1[i])
		i++
	}
	for j < len(nums2) {
		temp = append(temp, nums2[j])
		j++
	}
	return temp
}

// TestMergeSorts 归并排序
func TestMergeSorts(t *testing.T) {
	var nums = []int{2, 3, 5, 1, 6}
	merge_sort(nums, 0, len(nums)-1)
	fmt.Printf("归并排序结果为%v\n", nums)
}

func merge_sort(nums []int, left, right int) {
	if left >= right {
		//fmt.Printf("left %v >= right %v 直接返回\n", left, right)
		return
	}
	mid := (right + left) >> 1
	fmt.Printf("左侧处理 %v -- %v的数据\n", left, mid)
	merge_sort(nums, left, mid)
	fmt.Printf("右侧处理 %v -- %v的数据\n", mid+1, right)
	merge_sort(nums, mid+1, right)
	fmt.Printf("合并 %v -- %v的数据,中间点为%v\n", left, right, mid)
	merges(nums, left, mid, right)
}

func merges(nums []int, left, mid, right int) {
	//fmt.Printf("nums is %v ,left is %v mid is %v right is %v\n", nums, left, mid, right)
	//fmt.Printf("mid is %v\n", mid)
	//fmt.Printf("left is %v\n", nums[left:mid+1])
	//fmt.Printf("right is %v\n", nums[mid+1:right+1])
	var i, j = left, mid + 1

	//fmt.Printf("i is %v,j is %v\n", i, j)
	var temp []int
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[j])
			j++
		}
	}
	for i <= mid {
		temp = append(temp, nums[i])
		i++
	}
	for j <= right {
		temp = append(temp, nums[j])
		j++
	}
	//fmt.Printf("temp is %v\n", temp)
	for p, v := range temp {
		nums[left+p] = v
	}
}
