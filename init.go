package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var b int32
	a = 15
	b = b + 1
	a = a + a
	fmt.Print(a)
	fmt.Printf("ffff")
	fmt.Print(Unit8FromInt(9999))
}
func Unit8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil //nil表示零值，没有赋值的意思，返回两个参数，一个是 uint8 ，一个是error；
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}
