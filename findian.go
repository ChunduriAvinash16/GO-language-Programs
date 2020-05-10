package main
import "fmt"
//import "strings"
import "bufio"
import "os"
func main(){
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
	s := scanner.Text()
	flag1:=0
	if s[0]=='i' || s[0]=='I'{
		if s[len(s)-1]=='n' || s[len(s)-1]=='N'{
			for i:=1;i<len(s)-1;i++{
				if s[i]=='a' || s[i]=='A'{
					flag1=1
					break
				}
			}
			if flag1==1{
				fmt.Println("Found!")
			}else{
				fmt.Println("Not Found!")
			}
		}else{
			fmt.Println("Not Found!")
		}
	}else{
		fmt.Println("Not Found!")
	}
}