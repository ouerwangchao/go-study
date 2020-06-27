/**
  @author：wangchao
  @date：2020/6/25
  @note：
**/
package list

import (
	"fmt"
	"testing"
)

type Node struct {
	data interface{} //节点值 interface可以传任意类型的值
	next *Node       //下一个节点的指针
}

type List struct {
	Size int   //链表长度
	head *Node //头结点
}

//初始化链表
func (l *List) InitList() (list *List) {
	list = new(List)
	list.Size = 0
	var node = *new(Node)
	list.head = node.InitNode()
	return
}

//初始化节点
func (n *Node) InitNode() (node *Node) {
	var InitNode = new(Node)
	InitNode.data = nil
	InitNode.next = nil
	return InitNode
}

func TestListAdd(t *testing.T) {
	var list = new(List)
	currList := list.InitList()
	TailAddNode(currList, "111")
	TailAddNode(currList, "222")
	TailAddNode(currList, "333")
	TailAddNode(currList, "444")
	//
	//MiddleAddNode(currList, 3, "555")
	//UpdateNodeValue(currList, 2, "666")
	t.Log(currList.head.data)
	t.Log(currList.head.next.data)
	t.Log(currList.head.next.next.data)
	t.Log(currList.head.next.next.next.data)
	t.Log(currList.head.next.next.next.next.data)
	//DeleteNode(currList, 1)
	t.Log(currList.Size)
	ListReverse(currList)
	t.Log(currList.head.data)
	t.Log(currList.head.next.data)
	t.Log(currList.head.next.next.data)
	t.Log(currList.head.next.next.next.data)
	t.Log(currList.head.next.next.next.next.data)
}

//todo 跳表

/**
链表反转
*/
func ListReverse(list *List) {
	size := list.Size
	if size <= 1 {
		return
	}
	var prev = new(Node) //前驱节点
	pointer := list.head //当前节点 头结点开始
	for pointer != nil { //当前指针指向nil时停止
		prev, pointer, pointer.next = pointer, pointer.next, prev
	}
	list.head.next = prev
}

/**
删除节点操作
*/
func DeleteNode(list *List, idx int) {
	size := list.Size
	if idx > size || idx <= 0 {
		fmt.Printf("无效的删除操作")
		return
	}
	pointer := list.head
	for i := 1; i < idx; i++ {
		pointer = pointer.next
	}
	pointer.next = pointer.next.next
	list.Size -= 1
}

/**
尾插法
*/
func TailAddNode(list *List, data interface{}) {
	currentNode := list.head
	size := list.Size
	for i := 0; i < size; i++ {
		currentNode = currentNode.next //根据链表长度移动到最后一个节点
	}
	//构建一个插入节点
	var insertNode = new(Node)
	insertNode = insertNode.InitNode()
	insertNode.data = data
	currentNode.next = insertNode
	list.Size += 1
}

/**
更新节点数据
*/
func UpdateNodeValue(list *List, idx int, data interface{}) {
	size := list.Size //链表长度
	if idx > size || idx <= 0 {
		fmt.Printf("无效的节点操作")
		return
	}
	pointer := list.head //头结点
	for i := 0; i < idx; i++ {
		pointer = pointer.next
	}
	pointer.data = data
	return
}

/**
中间插入法
*/
func MiddleAddNode(list *List, idx int, data interface{}) bool {
	size := list.Size
	if idx > size+1 {
		fmt.Printf("插入失败，超出链表最大长度")
		return false
	}
	//头结点
	pointer := list.head
	for i := 1; i < idx; i++ {
		pointer = pointer.next
	}
	//构造插入节点
	insertNode := new(Node)
	insertNode.next = pointer.next
	insertNode.data = data
	//将当前节点的指针指向带插节点
	pointer.next = insertNode
	//链表长度+1
	list.Size += 1
	return true
}
