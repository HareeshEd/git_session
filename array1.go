package main

import "fmt"

func main() {
	var arr [5]int
	//var ptr *[5]int // declaring a pointer array method(1)
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	//ptr = &arr  method(1)
	//var ptr = &arr //declaring a pointer array method(2)
	ptr := &arr // declaring a pointer array method(3)
	ptr[3] = 33
	ptr[4] = 25
	fmt.Println(*ptr)
}
