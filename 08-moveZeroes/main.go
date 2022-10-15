package main

import "fmt"

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}

	fmt.Println(nums)
}

func main() {
	nums := []int{12, 0, 5, 8, 0, 0, 21, 2}
	moveZeroes(nums)
}
