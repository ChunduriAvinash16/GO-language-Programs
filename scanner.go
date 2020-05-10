package main

import (
	"fmt"
  "bufio"
  "os"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    line := scanner.Text()
   fmt.Println("captured:",line)
}