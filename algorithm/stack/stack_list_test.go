package stack

import "testing"

/**
切片实现栈操作
栈：后进先出
*/
func TestStack(t *testing.T) {
	var stack = make([]int, 0)
	//压栈
	stack = append(stack, 2)
	t.Log(stack)
	//出栈
	v := stack[len(stack)-1]     //返回值
	stack = stack[:len(stack)-1] //删除出栈元素
	t.Log(v)
	t.Log(stack)
}

/**
切片实现队列操作
队列：先进先出
*/
func TestQueue(t *testing.T) {
	var queue = make([]int, 0)
	//进队
	queue = append(queue, 2)
	//出队
	v := queue[0]
	t.Log(v)
	queue = queue[1:]
	t.Log(queue)
}
