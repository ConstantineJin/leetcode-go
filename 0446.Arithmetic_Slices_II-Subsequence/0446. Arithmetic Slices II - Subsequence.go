package main

func numberOfArithmeticSlices(nums []int) (ans int) {
	var f = make([]map[int]int, len(nums)) // f[i][d]表示尾项为nums[i]，公差为d的弱等差子序列（至少有两个元素的等差子序列）的个数
	for i, x := range nums {
		f[i] = map[int]int{}
		for j, y := range nums[:i] {
			var d = x - y
			ans += f[j][d]
			f[i][d] += f[j][d] + 1
		}
	}
	return
}
