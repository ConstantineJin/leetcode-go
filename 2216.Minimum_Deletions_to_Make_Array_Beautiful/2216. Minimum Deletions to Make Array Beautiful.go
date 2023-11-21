package main

func minDeletion(nums []int) (ans int) {
	n, even := len(nums), true // even记录删除操作后当前下标是否为偶数
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] && even { // 删除当前元素
			ans++
		} else {
			even = !even
		}
	}
	if (n-ans)%2 == 1 { // 删除元素后长度为奇数
		return ans + 1 // 删除最后一个元素
	}
	return ans
}
