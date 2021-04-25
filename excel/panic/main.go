package main

import "fmt"

func main() {
	hello()
}

func hello() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	panic("test panic")
}
