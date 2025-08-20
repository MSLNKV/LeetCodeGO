package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
	nums = []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
	nums = []int{3, 3}
	fmt.Println(twoSum(nums, 6))
}
func twoSum(nums []int, target int) []int {
	var answer []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				answer = append(answer, i, j)
				break
			}
		}
	}
	return answer
}
