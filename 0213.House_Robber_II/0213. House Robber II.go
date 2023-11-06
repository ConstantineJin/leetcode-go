package main

func rob(nums []int) int {
	switch len(nums) {
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	default:
		var r func(arr []int) int
		r = func(arr []int) int {
			x, y := 0, 0
			for _, v := range arr {
				z := max(x+v, y)
				x, y = y, z
			}
			return y
		}
		return max(r(nums[:len(nums)-1]), r(nums[1:])) // 首尾不能同时打劫，则将问题退化为第198题的方法，在抢劫0~倒数第二间和1~最后一间之间选择最大值
	}
}
