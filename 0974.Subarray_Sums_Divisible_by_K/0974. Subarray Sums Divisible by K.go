package main

func subarraysDivByK(nums []int, k int) (ans int) {
	var prefixSum int
	remainder := make([]int, k)
	remainder[0] = 1
	for _, num := range nums {
		prefixSum = ((prefixSum+num)%k + k) % k
		ans += remainder[prefixSum]
		remainder[prefixSum]++
	}
	return
}
