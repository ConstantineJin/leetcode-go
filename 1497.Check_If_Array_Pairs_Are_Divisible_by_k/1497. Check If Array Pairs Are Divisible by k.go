package main

func canArrange(arr []int, k int) bool {
	groups := make([]int, k) // 将数组元素按照其对 k 的余数分组统计个数
	for _, v := range arr {
		groups[(v%k+k)%k]++ // 注意处理负数
	}
	if groups[0]&1 != 0 { // 整除 k 的元素个数为奇数，无法配对
		return false
	}
	for i := 1; i < (k+1)/2; i++ {
		if groups[i] != groups[k-i] {
			return false
		}
	}
	return true // 题目给定 arr 长度为偶数，如果前面都能匹配，那么当 k 为偶数时，mod k 为 k/2 的个数必然也为偶数，可以完成匹配
}
