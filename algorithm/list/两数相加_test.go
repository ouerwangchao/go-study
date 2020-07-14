package list

import "testing"

type listNode struct {
	Val  int
	Next *listNode
}

func TestAdd(t *testing.T) {
	var l1 = listNode{Val: 2, Next: &listNode{Val: 4, Next: &listNode{Val: 3, Next: nil}}}
	var l2 = listNode{Val: 5, Next: &listNode{Val: 6, Next: &listNode{Val: 4, Next: nil}}}
	res := addList(&l1, &l2)
	t.Log(res)
}

func addList(l1, l2 *listNode) listNode {
	var l = &listNode{Val: 0, Next: nil}
	for l1 != nil || l2 != nil {
		l = l.Next
		if l1.Val+l2.Val > 9 {
			l.Val = 1
		} else {
			if l1.Next.Val+l2.Next.Val > 9 {
				l.Next.Val  = l1.Val + l2.Val + 1
			} else {
				l.Next.Val = l1.Val + l2.Val
			}
		}
	}
	return *l
}
