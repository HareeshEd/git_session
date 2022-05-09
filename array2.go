package main

import "fmt"

func main() {
	arr1 := [5]int{1: 10, 2: 40}
	arr2 := [...]string{"7up", "pepsi", "cococola", "slice", "sprite"}
	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
	fmt.Println(arr1, "\n", arr2)
}
