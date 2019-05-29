package main

import "fmt"

func main() {
	// fmt.Println(testReturn())
	a, b, _ := testReturn()
	fmt.Println(a)
	fmt.Println(b)
}

func testReturn() (string, string, int) {
	return "3", "2", 1
}