package main

import "fmt"

func main() {
	var total int
	_, _ = fmt.Scanln(&total)

	for i:=0;i< total;i++ {
		var num, target int
		fmt.Scanln(&num,&target)
		nums := make([]int,num)
		for j:=0;j<num;j++ {
			fmt.Scan(&nums[j])
		}
		n := Get(nums,target)
		if n>= 0 {
			fmt.Println(n)
		}else {
			fmt.Println("Impossible")
		}
	}
}

func Get(nums []int,aim int) int{
	if len(nums) < 3 {
		return  -1
	}
	sum := 0
	var tmps []int
	for i :=0;i<len(nums); i++ {
		tmps = append(tmps,nums[i])
	}
	for i,j := 0,2;j<len(nums);i,j=i+1,j+1 {
		tmp := 0
		for k :=i;k<=j;k++ {
			if tmps[k] <0 {
				tmps[k] = 0
			}
			tmp += tmps[k]
		}
		if tmp < aim {
			for l :=i;l <= j;l++ {
				if nums[l] <0 && tmp <aim{
					tmps[l] = aim- (tmp - tmps[l])
					tmp = aim
				}
			}
		}
		if tmp < aim {
			return  -1
		}
	}
	for i:=0;i<len(nums);i ++ {
		sum += tmps[i]
	}
	return sum
}
