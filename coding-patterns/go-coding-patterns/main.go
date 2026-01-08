package main

import (
	"fmt"
	"main/two-pointers"
)

func main() {
	nums := []int{-5, -2, 3, 4, 6}
	target := 7

	result := two_pointers.PairSumSortedBruteForce(nums, target)

	fmt.Println(result)
}
