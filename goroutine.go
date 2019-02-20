package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 4; i++ {
			fmt.Println("子进程1")
			time.Sleep(1 * time.Second)
			runtime.Goexit()  //退出该进程
			runtime.Gosched() //让出时间片
			fmt.Println("子go程")

		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("主进程")
		time.Sleep(1 * time.Second)
	}
}
