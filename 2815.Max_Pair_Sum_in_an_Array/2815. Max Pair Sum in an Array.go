package main

import "sort"

func getLargestDigit(x int) (res int) {
	for ; x > 0; x /= 10 {
		res = max(res, x%10)
	}
	return
}

func maxSum(nums []int) int {
	var count [10][]int
	for _, num := range nums {
		largestDigit := getLargestDigit(num)
		count[largestDigit] = append(count[largestDigit], num)
	}
	ans := -1
	for _, arr := range count {
		if len(arr) < 2 {
			continue
		}
		sort.Ints(arr) // 也可以使用优先队列或两个变量一次线性遍历维护最大和次大值
		ans = max(ans, arr[len(arr)-1]+arr[len(arr)-2])
	}
	return ans
}
