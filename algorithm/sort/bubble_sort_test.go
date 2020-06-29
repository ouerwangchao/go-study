/**
  @author：wangchao
  @date：2020/6/29
  @note：冒泡排序
**/
package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	var arr []int
	arr = []int{3, 4, 5, 34, 231, 521, 1, 512, 2}
	SortedArr := BubbleSort(arr)
	t.Log(SortedArr)
}

func BubbleSort(array []int) []int {
	arrLen := len(array)
	println(arrLen)
	for i := 0; i < arrLen-1; i++ { //每一次外层循环后都可以确认一个最大值放在末尾
		exchange := false //假设没有发生数据交换
		for j := 0; j < arrLen-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				exchange = true //发生数据交换
			}
		}
		if !exchange {
			//没有数据交换即可退出循环，减少循环比较次数
			break
		}
	}
	return array
}
