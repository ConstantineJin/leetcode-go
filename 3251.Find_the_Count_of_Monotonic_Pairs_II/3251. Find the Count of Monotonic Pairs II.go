package main

import "slices"

const mod int = 1e9 + 7

func countOfPairs(nums []int) (ans int) {
	n := len(nums)
	mx := slices.Max(nums)
	f := make([][]int, n) // f[i][j] 表示下标 0 到 i 中的单调数组对的个数，且 arr1[i]=j
	for i := range f {
		f[i] = make([]int, mx+1)
	}
	s := make([]int, mx+1)
	for j := 0; j <= nums[0]; j++ { // arr1[0] 可以取 [0,nums[0]]
		f[0][j] = 1
	}
	for i := 1; i < n; i++ {
		s[0] = f[i-1][0]
		for k := 1; k <= mx; k++ {
			s[k] = s[k-1] + f[i-1][k] // f[i-1] 的前缀和
		}
		for j := 0; j <= nums[i]; j++ { // 枚举 arr1[i]=j
			maxK := j + min(nums[i-1]-nums[i], 0) // 枚举 arr1[i−1]=k，有 k<=j 且 nums[i−1]−k>=nums[i]−j，整理得 k<=j+min(nums[i−1]−nums[i],0)
			if maxK >= 0 {
				f[i][j] = s[maxK] % mod // f[i-1][0]+f[i-1][1]+...+f[i-1][maxK]
			}
		}
	}
	for _, v := range f[n-1][:nums[n-1]+1] {
		ans += v
	}
	return ans % mod
}
