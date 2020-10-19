package main

import "fmt"

func main() {
	fmt.Println(find([]int{1, 2, 3, 2, 4}))
	fmt.Println(find2([]int{1, 2, 3, 2, 4}))

}
func find(nums []int) int {
	for _, v := range nums {
		fmt.Println(v)
		fmt.Println(nums[v])
		if nums[0] != nums[v] {
			nums[0], nums[nums[0]] = nums[nums[0]], nums[0]
		}
		fmt.Println(nums)
	}
	return nums[0]
}

func find2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[0] == nums[nums[0]] {
			return nums[0]
		}
		nums[0], nums[nums[0]] = nums[nums[0]], nums[0]
	}
	return nums[0]
}
