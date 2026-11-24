package main

import "fmt"

// ProductExceptSelf returns, for each index, the product of every other
// element — without division. Two sweeps: prefix products land in the
// output, then a running suffix product multiplies in from the right.
func ProductExceptSelf(nums []int) []int {
	out := make([]int, len(nums))
	prefix := 1
	for i, v := range nums {
		out[i] = prefix
		prefix *= v
	}
	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		out[i] *= suffix
		suffix *= nums[i]
	}
	return out
}

func main() {
	fmt.Println(ProductExceptSelf([]int{1, 2, 3, 4}))
}
