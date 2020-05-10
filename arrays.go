package main
import "fmt"
/*func main() {
	var x [5]int
	//for i:=0;i<len(x);i++{
	//	fmt.Println(x[i])
//	}
		for i,v:=range x{
			fmt.Println(v,i)
		}
}*/
func main() {
	x := [...]int {1, 2, 3, 4, 5}
	y := x[0:2]
	z := x[1:4]
	fmt.Print(len(y), cap(y), len(z), cap(z))
  }