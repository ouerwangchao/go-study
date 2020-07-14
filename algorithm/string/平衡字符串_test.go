package string

import (
	"fmt"
	"testing"
)

/*
在一个「平衡字符串」中，'L' 和 'R' 字符的数量是相同的。

给出一个平衡字符串 s，请你将它分割成尽可能多的平衡字符串。

返回可以通过分割得到的平衡字符串的最大数量。
*/
func TestBalanceStr(t *testing.T) {

	var str = "RLRRLLRLRL"
	fmt.Printf("v is %v s is %v\n", str[0], str[0:1])
	num := balanceSplit(str)
	fmt.Printf("%v分割可以得到的平衡字符串最大数量为：%v\n", str, num)

	num = commonBalance(str)
	fmt.Printf("%v方法2分割可以得到的平衡字符串最大数量为：%v\n", str, num)
}

func commonBalance(str string) int {
	var sum, num = 1, 0
	if str[:1] == "L" {
		sum = -1
	}
	l := len(str)
	for i := 1; i < l; i++ {
		if str[i:i+1] == "L" {
			sum--
		} else {
			sum++
		}
		if sum == 0 {
			num++
		}
	}
	return num
}

//todo 栈实现
func balanceSplit(s string) int {
	if len(s) == 0 {
		return 0
	}
	stack := []byte{s[0]}
	var res int
	for i := 1; i < len(s); i++ {
		if len(stack) > 0 && ((s[i] == 'L' && stack[len(stack)-1] == 'R') || (s[i] == 'R' && stack[len(stack)-1] == 'L')) {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				res++
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return res
}
