package main

import "fmt"

func main() {
	var (
		aa [2][2]int
		bb [2][2]int
		cc [2][2]int
	)
	fmt.Println("Enter the values: ")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&aa[i][j])
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&bb[i][j])
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			cc[i][j] = aa[i][j] + bb[i][j]
		}
	}
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
}
