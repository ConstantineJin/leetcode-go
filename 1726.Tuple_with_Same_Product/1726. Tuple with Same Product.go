package main

func tupleSameProduct(nums []int) (ans int) {
	products := make(map[int]int)
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			products[nums[i]*nums[j]]++
		}
	}
	for _, product := range products {
		ans += product * (product - 1) * 4
	}
	return
}
