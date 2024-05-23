package main

// 创建一个哈希表，key是元素值，value是元素在nums中的下标集合
// 假设nums的一个等值子数组的元素下标从pos[left]到pos[right]，在删除前有pos[right]-pos[left]+1个元素
// 其中有right-left+1个元素是相同的，那么需要删除pos[right]-pos[left]-(right-left)个元素
// 如果该式大于k则不断左移left直到其不超过k，此时使用right-left+1尝试更新答案
func longestEqualSubarray(nums []int, k int) (ans int) {
	posMap := make(map[int][]int)
	for i, num := range nums {
		posMap[num] = append(posMap[num], i-len(posMap[num])) // 实际的value是pos[i]-i
	}
	for _, pos := range posMap {
		if len(pos) <= ans {
			continue
		}
		var left int
		for right, p := range pos {
			for ; p-pos[left] > k; left++ { // 那么与k做比较的数就是pos[right]-pos[left]
			}
			ans = max(ans, right-left+1)
		}
	}
	return
}
