package main

import "fmt"
func Print(x *[3]int)  {
	//fmt.Println("Hello , world")
	(*x)[0] = (*x)[0] + 1
}
func main() {
	//a := [3]int{1,2,3}
	var a [3]int
	//b:=3
	fmt.Println(a)
	Print(&a)
	fmt.Println(a)
}