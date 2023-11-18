package main

func maximumSum(nums []int) int {
	ans, mp := -1, make(map[int]int) // mp的key为数位和，value为数位和=key的最大的nums[i]
	for _, num := range nums {
		sum := 0 // 当前元素的数位和
		for x := num; x > 0; x /= 10 {
			sum += x % 10
		}
		if _, ok := mp[sum]; ok { // 如果已经存在数位和为sum的元素
			ans = max(ans, mp[sum]+num) // 酌情更新
		}
		mp[sum] = max(mp[sum], num) // 保证mp[sum]的value是数位和为sum的nums[i]最大值
	}
	return ans
}
