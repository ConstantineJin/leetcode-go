package main

import "sort"

var palindromes = make([]int, 0, 109999)

func init() { // 按顺序从小到大生成所有回文数
	for base := 1; base <= 10000; base *= 10 {
		for i := base; i < base*10; i++ {
			x := i
			for t := i / 10; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			palindromes = append(palindromes, x)
		}
		if base <= 1000 {
			for i := base; i < base*10; i++ {
				x := i
				for t := i; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				palindromes = append(palindromes, x)
			}
		}
	}
	palindromes = append(palindromes, 1e9+1) // 哨兵，防止下标越界
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func minimumCost(nums []int) int64 {
	sort.Ints(nums)
	var cost = func(i int) (res int64) { // 返回所有 nums[i] 变成 pal[i] 的总代价
		var target = palindromes[i]
		for _, num := range nums {
			res += int64(abs(num - target))
		}
		return
	}
	var n = len(nums)
	var i = sort.SearchInts(palindromes, nums[(n-1)/2])
	if palindromes[i] <= nums[n/2] { // 回文数在中位数范围内
		return cost(i)
	}
	return min(cost(i-1), cost(i)) // 枚举离中位数最近的两个回文数 pal[i-1] 和 pal[i]
}
