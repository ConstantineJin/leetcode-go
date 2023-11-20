package main

func findShortestSubArray(nums []int) int {
	deg, cnt, first, last, ans := 0, make(map[int]int), make(map[int]int), make(map[int]int), len(nums) // deg为数组的度，cnt的key为数组元素，value为该元素出现次数，first与last分别记录数组元素首次及末次出现的位置
	for i, num := range nums {
		if _, ok := cnt[num]; !ok {
			first[num] = i
		}
		last[num] = i
		cnt[num]++
	}
	for _, v := range cnt {
		deg = max(deg, v) // 求数组的度
	}
	for i, v := range cnt {
		if v == deg { // 该元素是该数组的众数（之一）
			ans = min(ans, last[i]-first[i]+1)
		}
	}
	return ans
}
