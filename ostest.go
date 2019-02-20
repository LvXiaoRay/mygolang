package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Create("D:/go_workspace/hello.txt")
	if err != nil {
		fmt.Println("创建失败！")
	}
	fp.WriteString("hello,\r\n这是第二行")
	defer fp.Close()
	fp2, err2 := os.OpenFile("D:/go_workspace/hello.txt", os.O_RDWR, 6)
	if err2 != nil {
		fmt.Println("打开失败！")
	}
	defer fp2.Close()
	count, _ := fp2.Seek(0, os.SEEK_END)
	fp2.WriteAt([]byte("hello ,this is two"), count)
	//r:=bufio.NewReader(fp2)
	fmt.Printf("%T", fp)
}
