package string

import (
	"fmt"
	"math"
	"testing"
)

func TestIntReverse(t *testing.T) {
	var num = -210300001
	res := intReverse(num)
	fmt.Printf("反转结果为：%v\n", res)
}

func intReverse(num int) int {
	var res = 0
	for num != 0 {
		now := num % 10 //余数为所需值
		if res*10 > math.MaxInt32 || (res*10 == math.MaxInt32 && now > 7) {
			res = 0
			break
		}
		if res*10 < math.MinInt32 || (res*10 == math.MinInt32 && now < -8) {
			res = 0
			break
		}
		res = res*10 + now
		num /= 10
	}
	return res
}
