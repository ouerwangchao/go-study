package stack

import (
	"fmt"
	"strings"
	"testing"
)

func TestValidBrackets(t *testing.T) {
	var str = "())("
	res := CheckBrackets(str)
	fmt.Printf("有效性：%v\n", res)
}

func CheckBrackets(str string) bool {
	//使用数组栈实现
	res := ArrayStack(str)
	return res
}

type Node struct {
	Val  string
	Next *Node
}

func ListStack(str string) bool {

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
