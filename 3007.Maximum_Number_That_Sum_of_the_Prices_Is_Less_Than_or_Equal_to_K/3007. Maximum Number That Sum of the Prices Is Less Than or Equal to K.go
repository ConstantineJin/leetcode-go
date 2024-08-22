package main

import "math/bits"

func findMaximumNumber(kk int64, x int) int64 {
	k := int(kk)
	var num, pre1 int                                  // 设 i 左边有 pre1 个编号是 x 的倍数且填了 1 的比特位
	for i := bits.Len(uint(k+1)<<x) - 1; i >= 0; i-- { // 从高到低构建 num 的每个比特位，设当前枚举到 num 的从低到高的第 i 个比特位（i 从 0 开始）。[1,(k+1)·2^x]的价值和至少是 k+1，故至多考虑[1,(k+1)·2^x-1]的价值和。因而 i 初始化为 (k+1)·2^x 的最高位
		cnt := pre1<<i + i/x<<i>>1 // 如果 num 在第 i 位上填 1，在 i 左侧会增加价值 pre1·2^i，在 i 右侧会增加 (i/x)·(2^(i-1))
		if cnt > k {               // 如果 cnt 不超过 k 就填 1，由于是从高位到低位枚举，因此会使答案尽可能大
			continue
		}
		k -= cnt
		num |= 1 << i // 将第 i 位填为 1
		if (i+1)%x == 0 {
			pre1++
		}
	}
	return int64(num - 1)
}
