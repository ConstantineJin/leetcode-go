package main

func tupleSameProduct(nums []int) (res int) {
	products := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			products[nums[i]*nums[j]]++
		}
	}
	for _, v := range products {
		res += v * (v - 1) * 4
	}
	return
}
