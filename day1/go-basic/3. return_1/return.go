package main

import "fmt"

func main() {
	fmt.Println(testReturn())
}

func testReturn() (string, string) {
	return "3", "2"
}