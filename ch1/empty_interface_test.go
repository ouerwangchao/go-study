package ch1

import (
	"fmt"
	"testing"
)

func DoSomeThing(p interface{}) {
	//if i,ok := p.(int);ok{
	//	fmt.Println("int",i)
	//	return
	//}
	//if i,ok:= p.(string);ok{
	//	fmt.Println("string ",i)
	//	return
	//}
	//fmt.Println("no match")

	switch i := p.(type) {
	case int:
		fmt.Println("int", i)

	case string:
		fmt.Println("string", i)

	default:
		fmt.Println("no match")
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomeThing(1)
	DoSomeThing("waa")
	DoSomeThing([...]int{1, 2, 3})
	DoSomeThing([]int{})
}
