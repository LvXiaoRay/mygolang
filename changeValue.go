package main

import "fmt"

func main() {
	b := changevaleue(3)
	fmt.Println(b)
}
func changevaleue(a int) int {
	a = 100
	return a
}
