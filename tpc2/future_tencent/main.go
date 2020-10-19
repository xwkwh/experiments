package main

import "fmt"

func main() {
	//var total int
	//fmt.Scanln(&total)
	//
	//for i:=0;i< total;i++ {
	//	var num int
	//	fmt.Scanln(&num)
	//	if num > 0 {
	//		fmt.Println("You are the future of Tencent!")
	//	}else {
	//		fmt.Println("Good luck and Enjoy TPC!")
	//	}
	//}

	a := []int{1,2,4}
	c := a
	c[2] = 3
	fmt.Print(a)
	fmt.Print(c)
}
