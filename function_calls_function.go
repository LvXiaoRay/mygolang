package main

var aa string

func main() {
	aa = "G"
	print(aa)
	f1()
}
func f1() {
	aa := "0"
	print(aa)
	f2()
}
func f2() {
	print(aa)
}
