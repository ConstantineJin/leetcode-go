package main

func maximumLength(nums []int, k int) (ans int) {
	f := make([][]int, k) // f[y][x] 表示最后两项模 k 分别为 y 和 x 的子序列的长度
	for i := range f {
		f[i] = make([]int, k)
	}
	for _, x := range nums {
		x %= k
		for y, fxy := range f[x] {
			f[y][x] = fxy + 1
			ans = max(ans, f[y][x])
		}
	}
	return
}
