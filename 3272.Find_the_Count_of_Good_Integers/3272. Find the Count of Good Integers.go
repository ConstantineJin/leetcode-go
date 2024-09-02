package main

import (
	"math"
	"slices"
	"strconv"
)

var factorial [11]int

func init() {
	factorial[0] = 1
	for i := 1; i <= 10; i++ {
		factorial[i] = factorial[i-1] * i
	}
}

func countGoodIntegers(n, k int) (ans int64) {
	vis := make(map[string]bool)
	base := int(math.Pow10((n - 1) / 2))
	for i := base; i < base*10; i++ { // 枚举回文数左半边
		x := i
		t := i
		if n%2 > 0 {
			t /= 10
		}
		for ; t > 0; t /= 10 {
			x = x*10 + t%10
		}
		if x%k > 0 { // 回文数不能被 k 整除
			continue
		}
		bs := []byte(strconv.Itoa(x))
		slices.Sort(bs)
		s := string(bs)
		if vis[s] { // 避免重复统计
			continue
		}
		vis[s] = true
		var cnt [10]int
		for _, c := range bs {
			cnt[c-'0']++
		}
		res := (n - cnt[0]) * factorial[n-1] // 排列个数为 ((n-cnt0)(n-1)!)/(cnt0!*cnt1!*...*cnt9!)
		for _, c := range cnt {
			res /= factorial[c]
		}
		ans += int64(res)
	}
	return
}
