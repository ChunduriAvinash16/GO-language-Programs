package main
import (
	"fmt"
	"encoding/json"
)
type Person struct{
	name string
	address string
}
func main()  {
	var P1 Person
	fmt.Println("Enter your name")
	fmt.Scan(&P1.name)
	fmt.Println("Enter your Address")
	fmt.Scan(&P1.address)
	fmt.Println(P1)
	barr,err := json.Marshal(P1)
		if err!=nil{
			fmt.Println("ERRor")
		}
	fmt.Println(barr)
}