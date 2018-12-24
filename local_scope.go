package main

var a = "G"

func main() {
	b()
	d()
	b()

}
func b() { print(a) }

func d() {
	a := "0"
	print(a)
}
