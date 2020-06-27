package ch1

import (
	"fmt"
	"testing"
)

func TestChangeParam(t *testing.T)  {
	t.Log(Sum(1,2,3,4,5))
}

func Sum(ops ...int) int  {
	ret := 0
	for _,v := range ops{
		ret += v
	}
	return ret
}

func TestDefer(t *testing.T){
	defer fmt.Print(Sum(1,2,3))
	fmt.Println(12121)
	panic("432424")
}
