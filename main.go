package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	aaa()
	bbb()
	fmt.Println("Hello World")
	fmt.Println("Complete!!!!!")
}

func aaa() {
	fmt.Println("aaa")
}

func bbb() {
	bbb := "bbb"
	fmt.Println(bbb)
}
