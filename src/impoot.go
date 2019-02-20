package main

import (
	"Test"
	"fmt"
) //fuck,,,,,,引用的包名必须要 首字母大写，

func main() {
	Test.Test4() //方法也要首字母大写
	var array1 = []int{2, 3, 4, 5}
	var array2 = array1 //分配的是同一个空间,
	var bb = 99
	var c = bb
	c = 100
	fmt.Printf("%d\n", bb)
	fmt.Printf("%d\n", c)
	fmt.Printf("%p\n", array1)
	fmt.Printf("%p\n", array2)
	sum1 := append(array1, 3, 4, 5)
	sum2 := make([]int, 5, 6) //第三个参数表示容量，只有在重新切片时才会用到,sum2的长度为5；
	sum3 := array1[0:2:2]     //表示0-2的切片，后面的2表示分配给sum3的容量，使用方法---》切片名[起始位置，结束位置+1，容量]
	fmt.Println("ffffff")
	fmt.Println(sum3)
	sum2[1] = 0
	b := sum2[:cap(sum2)] //这里b对sum2进行切片，切片之后，b长度为6，切片后对应的是同一个空间，b内容改变，sum2也改变
	fmt.Println(sum1)
	fmt.Println(sum2)
	fmt.Println(len(sum2))
	fmt.Println(b)
	fmt.Printf("%T", sum1)
}
