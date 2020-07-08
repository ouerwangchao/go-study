package array

import (
	"testing"
)

var Array = []int{1, 2, 6, 2, 4, 5}

/**
查找数组中缺失和重复的元素
*/
func TestFindDupMiss(t *testing.T) {
	//todo 采用将将元素值-1作为数组下标并将元素值*-1的方式，重复的元素经过两次*-1操作，依旧是正值
	size := len(Array)
	var dup, miss int
	for i := 0; i < size; i++ {
		var idx int
		if Array[i] < 0{
			idx = -Array[i] -1
		}else{
			idx = Array[i] - 1
		}
		if Array[idx] < 0 {
			//表示已经反转过一次
			dup = Array[idx] * -1
		} else {
			Array[idx] *= -1
		}
	}

	for j := 0; j < size; j++ {
		if Array[j] > 0 {
			miss = j + 1
		}
	}
	t.Log(dup, miss, Array)
}
