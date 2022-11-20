package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	aaa()
	bbb()
	ccc()
	fmt.Println("Complete!!!!!")
}

func aaa() {
	fmt.Println("aaa")
}

func bbb() {
	bbb := "bbb"
	fmt.Println(bbb)
}

func ccc() {
	fmt.Println("ccc")
}

func ccc() {
	ccc := "ccc"
	fmt.Println(ccc)
}
