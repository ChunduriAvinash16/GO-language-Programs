package main
import "fmt"
func main()  {
	var i int
	fmt.Scan(&i)
	switch  {
	case i%2==0:fmt.Printf("Even")
	case i%2!=0:fmt.Printf("Odd")
	}
	
}