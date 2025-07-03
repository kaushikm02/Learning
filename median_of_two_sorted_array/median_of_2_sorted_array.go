package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (m+n+1)/2 - partitionX

		maxLeftX := math.Inf(-1)
		if partitionX != 0 {
			maxLeftX = float64(nums1[partitionX-1])
		}

		minRightX := math.Inf(1)
		if partitionX != m {
			minRightX = float64(nums1[partitionX])
		}

		maxLeftY := math.Inf(-1)
		if partitionY != 0 {
			maxLeftY = float64(nums2[partitionY-1])
		}

		minRightY := math.Inf(1)
		if partitionY != n {
			minRightY = float64(nums2[partitionY])
		}

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (m+n)%2 == 0 {
				return (math.Max(maxLeftX, maxLeftY) + math.Min(minRightX, minRightY)) / 2
			} else {
				return math.Max(maxLeftX, maxLeftY)
			}
		} else if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	panic("Input arrays are not sorted correctly")
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))    // Output: 2.0
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // Output: 2.5
}
