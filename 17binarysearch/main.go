package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Exploring Binary search in golang!!")
	arr := []int{}
	arr = append(arr, 4, 56, 7, 56, 89, 100, 2)
	sort.Ints(arr)
	fmt.Println(BinarySearch(arr, 0, len(arr)-1, 89))
	fmt.Println(BinarySearch(arr, 0, len(arr)-1, 678))
}

func BinarySearch(arr []int, low int, high int, val int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if arr[mid] == val {
		return mid
	} else if val < arr[mid] {
		return BinarySearch(arr, low, mid, val)
	} else {
		return BinarySearch(arr, mid+1, high, val)
	}
}
