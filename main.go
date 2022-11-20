package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	aaa()
	bbb()
	fmt.Println("Complete!!!!!")
}

func aaa() {
	fmt.Println("aaa")
}

func bbb() {
	bbb := "bbb"
	fmt.Println(bbb)
}
