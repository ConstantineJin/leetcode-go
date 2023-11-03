package main

func longestConsecutive(nums []int) (res int) {
	mp := make(map[int]bool)
	for _, num := range nums {
		mp[num] = true
	}
	for num := range mp {
		if !mp[num-1] { // 前一个数不存在
			cur, length := num, 1
			for mp[cur+1] { // 后一个数存在时
				cur++
				length++
			}
			res = max(res, length)
		}
	}
	return
}
