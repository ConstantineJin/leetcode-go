package main

import (
	"math"
	"slices"
	"sort"
)

// 树状数组模板（维护前缀最大值）
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}

func (f fenwick) preMax(i int) int {
	mx := math.MinInt
	for ; i > 0; i &= i - 1 {
		mx = max(mx, f[i])
	}
	return mx
}

func maxBalancedSubsequenceSum(nums []int) int64 {
	// 离散化 nums[i]-i
	b := slices.Clone(nums)
	for i := range b {
		b[i] -= i
	}
	slices.Sort(b)
	b = slices.Compact(b) // 去重

	// 初始化树状数组
	t := make(fenwick, len(b)+1)
	for i := range t {
		t[i] = math.MinInt
	}

	ans := math.MinInt
	for i, x := range nums {
		j := sort.SearchInts(b, x-i) + 1 // nums[i]-i 离散化后的值（从 1 开始）
		f := max(t.preMax(j), 0) + x
		ans = max(ans, f)
		t.update(j, f)
	}
	return int64(ans)
}
