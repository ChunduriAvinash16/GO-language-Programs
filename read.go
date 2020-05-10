package main
import (
	"fmt"
	"strings"
)
import "io/ioutil"
type Name struct{
	fname string
	lname string
}
func main(){
	var n [10]Name
	data,_:=ioutil.ReadFile("text.txt")
	//barr:=make([]byte,120)
	s:=string(data)
	//fmt.Println(s)
	s1:=strings.SplitAfter(s, "\n")
	for i:=0;i<len(s1);i++{
		modified:=strings.SplitAfter(string(s1[i])," ")
		n[i].fname=string(modified[0])
		n[i].lname=string(modified[1])
		fmt.Println(n[i])
	}
	

	
	
}

