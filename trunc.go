package main

import "fmt"
func main()  {
	var fnum float64
	fmt.Println("Enter the floating point number")
	fmt.Scan(&fnum)
	fmt.Println("Integer value is:")
	fmt.Printf("%d",int(fnum))
	
}