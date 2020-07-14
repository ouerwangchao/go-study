package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestZSatr(t *testing.T) {
	var str = "LEETCODEISHIRING"
	var row = 4
	res := strChange(str, row)
	fmt.Printf("%v按%v行做Z形变换后为：%v\n", str, row, res)
}

func strChange(str string, row int) string {
	if len(str) <= 2 || row == 1{
		return str
	}
	//col := int(math.Ceil(float64(len(str))/float64(row+row-2)) * float64(row-1))
	//fmt.Printf("%v行%v列\n", row, col)
	var arr = make([]string, row)
	//fmt.Printf("arr is %v\n", arr)
	var strArr = strings.Split(str, "")
	var curr = 0
	var direct = -1
	for _, v := range strArr { //todo 先赋值、再判断方向、最后移动
		arr[curr] += v
		if curr == 0 || curr == row-1 {
			direct = -direct
		}
		curr += direct
	}
	//fmt.Printf("arr is %v\n", arr)
	var res = ""
	for _, v := range arr {
		res += v
	}
	return res
}
