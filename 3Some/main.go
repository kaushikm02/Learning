package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates for the first number
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for second number
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for third number
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1 -1 2] [-1 0 1]]
	fmt.Println(threeSum([]int{0, 1, 1}))             // []
	fmt.Println(threeSum([]int{0, 0, 0}))             // [[0 0 0]]
}
