package main

func subarraySum(nums []int, k int) (res int) {
	pre, mp := 0, map[int]int{0: 1} // pre为前缀和，mp的key为和，value为出现次数
	for i := 0; i < len(nums); i++ {
		pre += nums[i]              // 更新前缀和
		if _, ok := mp[pre-k]; ok { // mp中已有键pre-k，亦即从该处到当前所遍历至构成的子数字之和为k
			res += mp[pre-k] // 结果添加对应的出现次数
		}
		mp[pre]++ // 更新mp
	}
	return
}
