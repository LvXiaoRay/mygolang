package main

import "fmt"

func main() {
	var maptt = make(map[string]int)  //map[string]int表示，string是key，int是value，key，value支持复合类型
	var maptt2 = make(map[int][2]int) //func make(t Type, size ...IntegerType) Type,int是key，【2】int是value
	maptt2[111] = [2]int{1, 3}
	maptt["ceshi"] = 100
	fmt.Println(maptt)
	fmt.Println(maptt["ceshi"])
	var p *int
	p = new(int)
	var a = 10
	*p = a
	fmt.Println(p)
}
