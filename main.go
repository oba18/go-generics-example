package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	testBranch()
	testBranch2()
	testBranch3()
	fmt.Println("Complete!!!!!")
}

func testBranch() {
	fmt.Println("Good Morning")
}

func testBranch2() string {
	s := "Good Evening"
	return s
}

func testBranch3() []int {
	is := []int{1, 2, 3}
	return is
}
