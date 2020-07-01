/**
  @author：wangchao
  @date：2020/7/1
  @note：快速排序
**/
package sort

import (
	"testing"
)

func TestSpeedSort(t *testing.T) {
	var array = []int{1, 3, 2, 54, 223, 666, 2, 23, 12, 53}
	res := SppedSortOne(array, 0, len(array)-1)
	t.Log(res)
	//res = SpeedSortTwo(array)
	//t.Log(res)
}

/*
快排思路，选取一个中心轴，将小于中心轴额放在左边，大的放在右边。重复以上步骤
*/
func SppedSortOne(array []int, left int, right int) []int {
	if left > right {
		return array
	}
	middle := array[left] //选定第一个元素作为中心轴
	l, r := left, right
	for {
		for array[r] >= middle && l < r {
			//右侧值大于等于中心轴时，不交换数据，向左移动一位
			r--
		}
		for array[l] <= middle && l < r {
			//左侧值小于等于中心轴时，不交换数据，向右移动一位
			l++
		}

		if l >= r {
			//将中心轴的值移动到当前位置
			array[left] = array[l]
			array[l] = middle
			break
		}
		//其他情况下，左侧和右侧值互换
		array[l], array[r] = array[r], array[l]

	}
	SppedSortOne(array, left, l-1)
	SppedSortOne(array, l+1, right)
	return array
}

//错误操作，待研究
func SppedSortOnes(array []int, left int, right int) []int {
	if left > right {
		return array
	}
	middle := array[left] //选定第一个元素作为中心轴
	l, r := left, right
	for {
		if array[r] < middle{  //1, 3, 2, 54, 223, 666, 2, 23, 12, 53
			//右侧值大于等于中心轴时，不交换数据，向左移动一位
			array[l] ,array[r] = array[r] , array[l]
			l++
		} else {
			r--
		}
		if array[l] > middle{
			//左侧值小于等于中心轴时，不交换数据，向右移动一位
			array[l] ,array[r] = array[r] , array[l]
			r--
		} else {
			l++
		}

		if l >= r {
			//将中心轴的值移动到当前位置
			array[left] = array[l]
			array[l] = middle
			break
		}
	}
	println(l,r)
	//SppedSortOne(array, left, l-1)
	//SppedSortOne(array, l+1, right)
	return array
}

//func SpeedSortTwo(array []int) []int {
//	if len(array) <= 1 {
//		return array
//	}
//	middle := array[0]
//	left, right := 1, len(array)-1
//	if left == right && middle > array[left] {
//		array[0], array[left] = array[left], array[0]
//		return array
//	}
//	for {
//		if array[right] < middle && right > left {
//			array[right], array[left] = array[left], array[right]
//			right--
//		}
//		if array[left] > middle && right > left {
//			array[right], array[left] = array[left], array[right]
//			left++
//		}
//		if left >= right {
//			array[left] = middle
//			break
//		}
//	}
//	SpeedSortTwo(array[:left])
//	SpeedSortTwo(array[left+1:])
//	return array
//}
