package main

import "fmt"

func main() {

LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf(" i is :%d ,and j is %d \n", i, j)
		}

	}

	a := 1
	b := 9
	goto TARGET

TARGET:
	b += a
	fmt.Printf("a is %v  b is %v", a, b)

	var i1 int
	var f1 float32
	i1, _, f1 = ThreeValues()
	fmt.Printf("%d ,%f", i1, f1)
}
func ThreeValues() (int, int, float32) {
	return 5, 6, 7.5

}
