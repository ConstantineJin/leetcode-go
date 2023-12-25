package main

// 单调栈
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var mp, st, ans = make(map[int]int), make([]int, 0), make([]int, len(nums1)) // 元素各不相同，故可以用哈希表
	for i := len(nums2) - 1; i >= 0; i-- {                                       // 倒序遍历nums2
		var num = nums2[i]
		for len(st) > 0 && num >= st[len(st)-1] { // 当栈不为空且当前元素不小于栈顶元素时
			st = st[:len(st)-1] // 出栈
		}
		if len(st) > 0 { // 如果栈不为空
			mp[num] = st[len(st)-1] // 下一个更大元素就是栈顶元素
		} else {
			mp[num] = -1 // 不存在下一个更大元素，保存-1
		}
		st = append(st, num) // 当前元素入栈
	}
	for i, num := range nums1 {
		ans[i] = mp[num]
	}
	return ans
}
