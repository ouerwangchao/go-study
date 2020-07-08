package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestChildStr(t *testing.T) {
	var str = "abcabcbb"
	res := GetChildStr(str)
	fmt.Printf("%v的最长子串为%v\n", str, res)
}

func GetChildStr(str string) string {
	length := len(str)
	if length <= 1 {
		return str
	}
	arr := strings.Split(str, "")
	res := ""
	maxLen := 0
	for i := 0; i < length; i++ {
		var m = make(map[string]int)
		if maxLen < len(arr[i:]) {
			var l = 0
			var s string
			for _, v := range arr[i:] {
				//判断v是否在m中
				_, exist := m[v]
				if exist {
					break
				} else {
					m[v] = 1
					l++
					s += v
				}
			}
			if l > maxLen {
				maxLen = l
				res = s
			}
		}
	}
	return res
}
