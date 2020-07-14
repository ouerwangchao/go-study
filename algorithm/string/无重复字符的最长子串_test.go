package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestChildStr(t *testing.T) {
	var str = "tabcabcbbkdjsha"
	res := GetChildStr(str)
	fmt.Printf("%v的最长子串为%v\n", str, res)
	result := GetChildStrTwo(str)
	fmt.Printf("%v的最长子串长度为%v\n", str, result)
	r := GetChildStrThree(str)
	fmt.Printf("%v的最长子串长度为%v\n", str, r)
}

//滑动窗口解法
func GetChildStrThree(str string) int {
	length := len(str)
	var m = make(map[byte]int)
	var rk, num = 0, 0
	for i := 0; i < length; i++ {
		if i != 0 { //窗口每次右移，需要删除窗口第一个字符
			delete(m, str[i-1])
		}
		for rk < length && m[str[rk]] == 0 { //窗口向右扩大  m[str[rk]] == 0表示当前字符未出现
			m[str[rk]]++
			rk++
		}
		num = GetMax(num, rk-i)
	}
	return num
}

func GetMax(num, compare int) int {
	if num > compare {
		return num
	}
	return compare
}

//双指针
func GetChildStrTwo(str string) int {
	length := len(str)
	if length <= 1 {
		return length
	}
	var res = length //默认为整个字符串长度
	left := 0
	arr := strings.Split(str, "")
	//Start:
	for left < length {
		for i := 0; i <= length-res; i++ {

			//判断array[i:res+i]中是否有重复
			result := CheckRepeatTwo(arr[i : res+i])
			//fmt.Printf("取%v个，查询范围为：%v,结果为：%v\n", res, arr[i:res+i], result)
			if result {
				goto LOOP
				//break Start
			}
		}
		left++
		res--
	}
LOOP:
	return res
}

func CheckRepeatTwo(arr []string) bool {
	res := true
START:
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == arr[i] {
				res = false
				break START
			}
		}
	}
	return res
}

func CheckRepeat(arr []string) bool {
	m := make(map[string]int)
	res := true
	for _, v := range arr {
		_, exist := m[v]
		if exist {
			res = false
			break
		} else {
			m[v]++
		}
	}
	return res
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
