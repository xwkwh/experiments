package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	var total int
	fmt.Scanln(&total)

	//var res []string

	for i:=0;i< total;i++ {
		var num int
		fmt.Scanln(&num)
		input,_,_:= inputReader.ReadLine()
		numstrs := strings.Split(string(input)," ")
		nums := make([]int,num)
		for k,v := range numstrs {
			vv, _ := strconv.Atoi(v)
			nums[k] = vv
		}

		if Is(nums) {
			//res = append(res,"Yes")
			fmt.Println("Yes")
		}else {
			//res = append(res, "No")
			fmt.Println("No")
		}
	}
	//for _,v := range res {
	//	fmt.Println(v)
	//}
	return
}

func Is(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	for i:=0;i<len(nums);i++ {
		for j := i+1;j<len(nums);j++ {
			if !Have(nums[i:j+1]){
				return false
			}
		}
	}

	return true
}

func Have(nums []int) bool {
	have := false
	for _, v := range nums {
		if v == 0 {
			return false
		}
		for k, vv := range nums {
			if vv%v != 0 {
				break
			} else {
				if k == len(nums)-1 {
					have = true
					return true
				}
			}
		}
	}
	return have
}
