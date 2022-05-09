package main

import "fmt"

func main() {
	var sum int
	var num int
	fmt.Print("enter the number:")
	fmt.Scan(&num)
	sum = num * (num + 1) / 2
	fmt.Println("Sum of n numbers : ", sum)
}
