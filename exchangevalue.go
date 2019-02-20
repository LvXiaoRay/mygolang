package main

import (
	"fmt"
)

func swap(a int, b int) {
	temp := a
	a = b
	b = temp
}
func main() {
	b := 20
	a := 10
	var c *int //此时指针c指向的是一个空地址
	c = new(int)
	c = &a
	swap(a, b) //不会改变a,b的值
	fmt.Println(a, b)
	apoint := [5]int{}
	var p *[5]int //定义一个数组指针
	p = &apoint
	fmt.Println(*p)
	//定义一个指针数组
	var p2 = [5]*int{} //这个是指针数组
	p2[0] = c
	p2[1] = &b
	d := [3]int{2, 3, 4}
	e := [3]int{2, 3, 4}
	f := [3]int{2, 3, 4}
	fmt.Println(*p2[0])
	fmt.Println(*p2[1])

	var arraypoint = [3]*[3]int{&d, &e, &f}
	fmt.Println(arraypoint)
	for i := 0; i < len(arraypoint); i++ {
		for j := 0; j < len(arraypoint[0]); j++ {
			fmt.Print(arraypoint[i][j])
			fmt.Print(" ")

		}
		fmt.Println()
	}

}
