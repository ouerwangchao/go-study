package stack

import (
	"fmt"
	"strings"
	"testing"
)

func TestValidBrackets(t *testing.T) {
	var str = "()()"
	res := CheckBracketsArrayStack(str)
	fmt.Printf("数组栈方式验证有效性：%v\n", res)
	result := CheckBracketsListStack(str)
	fmt.Printf("链表栈方式验证有效性：%v\n", result)
}

func CheckBracketsArrayStack(str string) bool {
	//使用数组栈实现
	res := ArrayStack(str)
	return res
}

func CheckBracketsListStack(str string) bool {
	//使用链表栈实现
	res := ListStack(str)
	return res
}

type Node struct {
	Val  string
	Next *Node
}
type List struct {
	Size int
	Head *Node
}

func ListStack(str string) bool {
	var listStack = new(List) //链表栈
	listStack.Size = 0
	listStack.Head = new(Node)
	strArr := strings.Split(str, "")
	for _, v := range strArr {
		if (v == ")" || v == "}" || v == "]") && listStack.Size == 0 {
			return false
		}
		if v == "(" || v == "{" || v == "[" {
			//入栈
			fmt.Printf("%v入栈\n", v)
			PushToStack(listStack, v)
		} else {
			//取栈内最后一个node值与v匹配，匹配成功则继续，匹配失败直接 return false
			stackVal := PopFormStack(listStack)
			fmt.Printf("出栈一个值%v与%v匹配\n", stackVal, v)
			if !(stackVal == "{" && v == "}") && !(stackVal == "(" && v == ")") && !(stackVal == "[" && v == "]") {
				return false
				break
			}
		}
	}
	if listStack.Size == 0 {
		return true
	}
	return false
}

//入栈
func PushToStack(l *List, str string) bool {
	var n = new(Node)
	n.Val = str
	n.Next = nil
	ListSize := l.Size
	var now = l.Head //头结点
	for i := 0; i < ListSize; i++ {
		now = now.Next
	}
	now.Next = n
	l.Size += 1
	return true
}

//出栈并返回尾节点值
func PopFormStack(l *List) string {
	listSize := l.Size
	var pointer = l.Head
	for i := 0; i < listSize-1; i++ {
		pointer = pointer.Next
	}
	val := pointer.Next.Val
	pointer.Next = nil
	l.Size -= 1
	return val
}

//数组栈
func ArrayStack(str string) bool {
	var arrStack = make([]string, 0) //数组栈
	strArr := strings.Split(str, "")
	for _, v := range strArr {
		if v == "{" || v == "(" || v == "[" {
			arrStack = append(arrStack, v)
		} else {
			if len(arrStack) == 0 && (v == "}" || v == ")" || v == "]") {
				return false
				break
			} else if arrStack[len(arrStack)-1] == "{" && v == "}" {
				arrStack = arrStack[:len(arrStack)-1]
			} else if arrStack[len(arrStack)-1] == "(" && v == ")" {
				arrStack = arrStack[:len(arrStack)-1]
			} else if arrStack[len(arrStack)-1] == "[" && v == "]" {
				arrStack = arrStack[:len(arrStack)-1]
			} else {
				return false
				break
			}
		}
	}
	if len(arrStack) == 0 {
		return true
	} else {
		return false
	}
}
