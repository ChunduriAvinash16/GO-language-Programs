package main

import "fmt"

func Swap(slice []int, i int) {
	if slice[i] > slice[i+1] {
		slice[i], slice[i+1] = slice[i+1], slice[i]
	}
}

func BubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			Swap(slice, j)
		}
	}
}

func main() {
	var num int
	fmt.Printf("Enter the number of integers you want to sort: ")
	fmt.Scan(&num)

	var slice []int

	fmt.Printf("Enter those %d numbers: ", num)
	for i := 0; i < num; i++ {
		var element int
		fmt.Scan(&element)
		slice = append(slice, element)
	}

	fmt.Println("Given slice: ", slice)
	BubbleSort(slice)
	fmt.Println("Sorted slice: ", slice)
}