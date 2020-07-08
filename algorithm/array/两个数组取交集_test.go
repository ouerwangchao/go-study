package array

import (
	"fmt"
	"testing"
)

func TestOccur(t *testing.T) {
	nums1 := make([]int, 0)
	nums1 = []int{1, 2, 3, 4, 5, 7, 5}
	nums2 := []int{2, 3, 5, 5, 5, 1, 6}
	res := getArrayOccur(nums1, nums2)
	fmt.Printf("交集为：%v", res)
}

func getArrayOccur(firstArray []int, secondArray []int) []int {
	fmt.Printf("数组1：%v\n", firstArray)
	fmt.Printf("数组2：%v\n", secondArray)
	var firstMap = make(map[int]int)
	for _, val := range firstArray {
		_, exist := firstMap[val]
		if exist {
			firstMap[val] += 1
		} else {
			firstMap[val] = 1
		}
	}
	var secondMap = make(map[int]int)
	for _, val := range secondArray {
		_, exist := secondMap[val]
		if exist {
			secondMap[val] += 1
		} else {
			secondMap[val] = 1
		}
	}
	fmt.Printf("map1：%v\n", firstMap)
	fmt.Printf("map2：%v\n", secondMap)

	result := make([]int, 0)
	for k, v := range firstMap {

		vv, exist := secondMap[k]
		if exist {
			//fmt.Printf("k:%v,v:%v,vv:%v\n", k, v, vv)
			min := 0
			if v < vv {
				min = v
			} else {
				min = vv
			}
			for i := 0; i < min; i++ {
				result = append(result, k)
			}
		}
	}
	return result
}
