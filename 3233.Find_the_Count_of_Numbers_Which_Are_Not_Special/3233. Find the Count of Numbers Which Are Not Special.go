package main

import "math"

const mx = 31622 // √1e9

var pi = [mx + 1]int{}

func init() { // 筛法预处理质数
	for i := 2; i <= mx; i++ {
		if pi[i] == 0 { // i 是质数
			pi[i] = pi[i-1] + 1
			for j := i * i; j <= mx; j += i {
				pi[j] = -1 // 标记 i 的倍数为合数
			}
		} else {
			pi[i] = pi[i-1]
		}
	}
}

func nonSpecialCount(l, r int) int { // 区间 [0,i] 内的特殊数字个数等于[0,floor(√i)]中质数的个数
	cntL, cntR := pi[int(math.Sqrt(float64(l-1)))], pi[int(math.Sqrt(float64(r)))]
	return r - l + 1 - (cntR - cntL)
}
