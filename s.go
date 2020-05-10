package main

import "fmt"
import "strings"
import "bufio"
import "os"

func main () {
	fmt.Printf("Please enter a string: ")

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.ToLower(strings.TrimSpace(s))

	if strings.Index(s, "i") == 0 && strings.Contains(s, "a") && strings.LastIndex(s, "n") == len(s)-1 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
		
}