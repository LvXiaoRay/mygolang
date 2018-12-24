package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	var ster string = "123456"
	an, err := strconv.Atoi(ster) //该函数返回两个参数；int和error;字符串转数字
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(an)

	str := "fuck you ass hole dame it"
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d  /", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Println("%d /", r)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f /", 100*rand.Float32())
	}
	str2 := strings.Fields(str)
	for _, val := range str2 {
		fmt.Printf("%s  ", val)
	}

}

var prompt = "Enter  a digit ,e.g.3" + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z,Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}
