package main

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k <= 1 {
		return
	}
	var product, i = 1, 0
	for j, num := range nums {
		product *= num
		for product >= k {
			product /= nums[i]
			i++
		}
		ans += j - i + 1
	}
	return
}
