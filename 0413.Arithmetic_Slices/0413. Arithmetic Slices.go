package main

func numberOfArithmeticSlices(nums []int) (ans int) {
	var n = len(nums)
	if n < 3 {
		return
	}
	var d, t = nums[0] - nums[1], 0 // d为当前公差，t为每次答案增加的次数
	for i := 2; i < n; i++ {
		if nums[i-1]-nums[i] == d { // 如果当前数与前一个数之差等于公差d
			t++
		} else {
			d, t = nums[i-1]-nums[i], 0 // 回到初始化
		}
		ans += t
	}
	return
}
