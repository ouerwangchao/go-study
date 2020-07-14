package string

import (
	"strconv"
	"testing"
)

func TestIntReverse(t *testing.T) {
	var num = -210300000
}

func intReverse(num int) int {
	var t = 1 //0表示正数
	n := num
	if num < 0 {
		t = -1
		n = -num
	}
	//将n反转
	s := strconv.Itoa(n)
	d :=1//判断当前字符是否有效




	return 1
}
