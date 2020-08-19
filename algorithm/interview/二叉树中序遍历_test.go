package interview

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initTree() *TreeNode {
	root := TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Right.Left = &TreeNode{7, nil, nil}
	root.Right.Right = &TreeNode{8, nil, nil}
	root.Left.Left = &TreeNode{4, nil, nil}
	root.Left.Right = &TreeNode{6, nil, nil}
	return &root
}

var res = make([]int, 0)

func TestMidTree(t *testing.T) {
	treeNode := initTree()
	listNode(treeNode)
	fmt.Printf("%v\n", res)
}

func TestA(t *testing.T) {
	t.Log("2121")
}

func listNode(root *TreeNode) {
	if root == nil {
		return
	}
	left := root.Left
	right := root.Right
	res = append(res, root.Val)
	listNode(left)
	listNode(right)
}
