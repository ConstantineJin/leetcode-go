package main

type fenwick []int // 树状数组，单点修改，区间查询

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func (f fenwick) query(l, r int) int {
	if l > r {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}

func countOfPeaks(nums []int, queries [][]int) (ans []int) {
	n := len(nums)
	f := make(fenwick, n-1)
	update := func(i, val int) {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			f.update(i, val)
		}
	}
	for i := 1; i < n-1; i++ {
		update(i, 1)
	}
	for _, query := range queries {
		if query[0] == 1 { // 操作 1 相当于计算从 l+1 到 r−1 的子数组的元素和
			ans = append(ans, f.query(query[1]+1, query[2]-1))
		} else {
			i := query[1]
			for j := max(i-1, 1); j <= min(i+1, n-2); j++ {
				update(j, -1) // 先把区间中的峰值元素从树状数组中去掉
			}
			nums[i] = query[2] // 修改
			for j := max(i-1, 1); j <= min(i+1, n-2); j++ {
				update(j, 1) // 再把区间中的峰值元素加入树状数组
			}
		}
	}
	return
}
