package main

func numberOfSubstrings(s string) (ans int) {
	n := len(s)
	var a []int // s 中 0 的下标集合
	for i, b := range s {
		if b == '0' {
			a = append(a, i)
		}
	}
	total1 := n - len(a)     // s 中 1 的数量
	a = append(a, n)         // 哨兵，方便边界处理
	for left, b := range s { // 枚举子串左端点 left
		if b == '1' {
			ans += a[0] - left // 不含 0 的子串数量
		}
		for k, p := range a[:len(a)-1] { // 枚举子串中 0 的个数，注意 a[len(a)-1] 为先前添加的哨兵
			cnt0 := k + 1           // 子串中 0 的个数
			if cnt0*cnt0 > total1 { // 内循环时间复杂度优化为 √n，整体时间复杂度 n√n
				break
			}
			q := a[k+1]           // 下一个 0 的下标
			cnt1 := p - left - k  // 子串中 1 的个数
			if cnt0*cnt0 > cnt1 { // 1 的数量不够多，需要补充 cnt0*cnt0-cnt1 个 1
				ans += max(q-p-(cnt0*cnt0-cnt1), 0)
			} else { // 1 的数量足够多，p,p+1,...,q-1都可以作为子串的右端点
				ans += q - p
			}
		}
		if b == '0' {
			a = a[1:] // 这个 0 后续不会再枚举到
		}
	}
	return
}
