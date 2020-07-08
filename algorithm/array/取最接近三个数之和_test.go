package array

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestDoublePointer(t *testing.T) {
	var target = 10
	var array = []int{-1, 0, 2, 3, 7, 6, 5, 7}
	sort.Ints(array)
	res, str := DoublePoint(array, target)
	fmt.Printf("与%v最接近的三个数之和为%v,由%v求和\n", target, res, str)
}

/**
从数组array中取出和值最接近target值的三个相邻元素
*/
func DoublePoint(array []int, target int) (sum int, str string) {
	len := len(array)
	if len < 3 {
		return 0, ""
	}
	var diff int //表示和值与target值的差值（绝对值）
	var types bool = false
	var res int
	var s string

	for i := 1; i < len-1; i++ {
		sum := array[i-1] + array[i] + array[i+1]
		temp := sum - target
		if temp < 0 {
			temp = -temp
		}
		if !types {
			diff = temp
			types = true
			res = sum
			s = strconv.Itoa(array[i-1]) + "-" + strconv.Itoa(array[i]) + "-" + strconv.Itoa(array[i+1])
		} else {
			if temp < diff {
				diff = temp
				res = sum
				s = strconv.Itoa(array[i-1]) + "-" + strconv.Itoa(array[i]) + "-" + strconv.Itoa(array[i+1])
			}
		}

	}
	return res, s
}
