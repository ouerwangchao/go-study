package interview

import (
	"fmt"
	"testing"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func TestBiranyTree(t *testing.T) {
	var arr = []int{2, 4, 3, 6, 8, 7, 5}
	res := biranyTree(arr, "init")
	fmt.Printf("%v\n", res)
}

func biranyTree(arr []int, t string) *Node {
	fmt.Printf("%v is %v\n", t, arr)
	var n Node
	n.Val = arr[len(arr)-1]
	if len(arr) <= 1 {
		return &n
	}
	var mid = 0
	//查询数组中最后一个比arr[len(arr)-1]小的索引值
	mid = getMidIdx(arr, arr[len(arr)-1])
	n.Left = biranyTree(arr[:mid+1], "left")
	n.Right = biranyTree(arr[mid+1:len(arr)-1], "right")
	//fmt.Printf("%v\n", n)
	return &n
}

func getMidIdx(arr []int, v int) int {
	for i, val := range arr {
		if v > val && v < arr[i+1] {
			return i
		}
	}
	return 0
}
