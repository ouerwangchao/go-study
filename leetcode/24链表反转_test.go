package leetcode

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestListReverse(t *testing.T) {
	var head = new(ListNode)
	res := ListReverse(head)
	t.Log(res)
}

func ListReverse(head *ListNode) *ListNode {
	var prev *ListNode //前驱节点
	println(prev)
	pointer := head          //当前节点 头结点开始
	for pointer != nil {     //当前指针指向nil时停止
		prev, pointer, pointer.Next = pointer, pointer.Next, prev
	}
	return prev
}
