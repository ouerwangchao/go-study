package string

import (
	"fmt"
	"strings"
)

var sentence = "i use riple pillow"
var searchWord = "tri"

func main() {
	num := isPrefixOfWord(sentence, searchWord)
	fmt.Printf("num:%v", num)
}

func isPrefixOfWord(sentence, searchWord string) int {
	arr := strings.Split(sentence, " ")
	length := len(arr)
	if length == 0 {
		return -1
	}
	for i := 0; i < length; i++ {
		str := arr[i]
		if len(searchWord) > len(str) {
			continue
		}
		ok := 0
		for j := 0; j < len(searchWord); j++ {
			if searchWord[j] != str[j] {
				ok = 1
				break
			}
		}
		if ok == 0 {
			return i + 1
		}
	}
	return -1
}
