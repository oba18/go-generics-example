package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	testBranch()
	testBranch2()
}

func testBranch() {
	fmt.Println("Good Morning")
}

func testBranch2() string {
	s := "Good Evening"
	return s
}
