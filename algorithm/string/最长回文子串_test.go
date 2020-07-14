package string

import (
	"fmt"
	"testing"
)

func TestChildString(t *testing.T) {
	var str = "babad"
	res := getChildStrThree(str)
	fmt.Printf("%v的最长回文子串为%v\n", str, res)
}

func getChildStrThree(str string) string {
	length := len(str)
	if length <= 1 {
		return ""
	}
	var m = make(map[byte]int)
	var num, res = 0, "" //回文子串长度和回文子串
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if m[str[j]] == 0 { //不存在
				m[str[j]]++
				res += str[j : j+1]
			} else {
				res += str[j : j+1]
				num = getMax(num, len(m)+1)
				delete(m, str[j])
				m[str[j]]++
			}
		}
	}

}

func getMax(num, compare int) int {
	if num > compare {
		return num
	}
	return compare
}
