package main

import (
	"fmt"
)

// Function for linear search in an array
func linearSearch(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i // Returns the index of the element if found
		}
	}
	return -1 // Returns -1 if the element is not found
}

// Function for binary search in an array (assumes the array is sorted)
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		// If the element is at the middle, return the index
		if arr[mid] == target {
			return mid
		}

		// If the target is less than the element in the middle, search in the left half
		if target < arr[mid] {
			right = mid - 1
		} else { // Otherwise, search in the right half
			left = mid + 1
		}
	}

	return -1 // Returns -1 if the element is not found
}

func main() {
	// Example array
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	// Linear Search
	linearIndex := linearSearch(arr, 12)
	fmt.Println("Linear Search: 12 found at index", linearIndex)

	// Binary Search (array needs to be sorted)
	binaryIndex := binarySearch(arr, 14)
	fmt.Println("Binary Search: 14 found at index", binaryIndex)
}
