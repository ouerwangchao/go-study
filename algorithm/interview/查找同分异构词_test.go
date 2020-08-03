package interview

import (
	"fmt"
	"sync"
	"testing"
)

func TestFindStr(t *testing.T) {
	var str = "abcsdedfasa"
	var aim = "asa"
	idx := findStr(str, aim)
	fmt.Printf("%v\n", idx)
}

func findStr(s, a string) int {
	if len(a) > len(s) || len(a) == 0 {
		return 0
	}
	var aimMap = make(map[int]int)
	for _, v := range a {
		aimMap[int(v)]++
	}
	i := 0
	var arr []int
	for i < len(s) {
		if len(arr) < len(a) {
			arr = append(arr, int(s[i]))
			aimMap[int(s[i])]--
			if aimMap[int(s[i])] == 0 {
				delete(aimMap, int(s[i]))
			}
		}
		arr = append(arr, int(s[i]))
		aimMap[int(s[i])]--
		if aimMap[int(s[i])] == 0 {
			delete(aimMap, int(s[i]))
		}
		aimMap[arr[0]]++
		if aimMap[arr[0]] == 0 {
			delete(aimMap, arr[0])
		}
		arr = arr[1:]
		if len(aimMap) == 0 {
			return i - len(a) + 1
		}
		i++
	}
	return 1
}

func TestGo(t *testing.T) {
	var wg sync.WaitGroup
	i := 0
	for i < 10 {
		i++
		wg.Add(1)
		go func(j int) {
			fmt.Printf("%v\n", j)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
