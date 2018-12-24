package main

import (
	"fmt"
	"os"
	"runtime"
)

var Pi float64

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("操作系统是：%s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("路径是:%s\n", path)
	var ss float64 = PI * 2
	fmt.Printf("%g", ss)
}
