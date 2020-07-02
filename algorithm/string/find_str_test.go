package string

import (
	"fmt"
	"testing"
)

/**
查找targetStr在longStr中出现的第一个位置
*/
func TestFindStr(t *testing.T) {
	var longStr = "abccd"
	var targetStr = "cd"
	res := findStr(longStr, targetStr)
	t.Log(res)
}

func findStr(longStr, targetStr string) int {
	var i, j int
	if len(targetStr) == 0 || len(targetStr) > len(longStr) {
		return -1
	}
	//todo 不用循环到len(longStr)-1,因为目标字符串有长度，剩余长度不够目标串长度时，一定不会匹配成功
	for i = 0; i < len(longStr)-len(targetStr)+1; i++ {
		for j = 0; j < len(targetStr); j++ {
			if longStr[i+j] != targetStr[j] {
				fmt.Printf("i:%v j:%v\n", i, j)
				break
			}
		}
		if j == len(targetStr) { //todo targetStr的每一个字符都正好匹配
			fmt.Printf("ok j:%v len:%v\n", j, len(targetStr))
			return i
		}
	}
	return -1
}
