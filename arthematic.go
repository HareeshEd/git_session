package main

import "fmt"

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func mul(x, y int) int {
	return x * y
}
func div(x, y int) int {
	return x / y
}
func main() {
	var num1, num2 int
	fmt.Println("Enter the First Number = ")
	fmt.Scanln(&num1)

	fmt.Println("Enter the Second Number = ")
	fmt.Scanln(&num2)

	fmt.Println("The Sum of num1 and num2  = ", add(num1, num2))
	fmt.Println("The sub of num1 and num2  = ", sub(num1, num2))
	fmt.Println("The mul of num1 and num2  = ", mul(num1, num2))
	fmt.Println("The div of num1 and num2  = ", div(num1, num2))
}
