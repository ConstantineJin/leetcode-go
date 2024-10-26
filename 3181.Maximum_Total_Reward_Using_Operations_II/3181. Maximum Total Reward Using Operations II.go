package main

import (
	"math/big"
	"slices"
	"sort"
)

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)                     // 排序，从小到大选
	rewardValues = slices.Compact(rewardValues) // 去重
	one := big.NewInt(1)
	f := big.NewInt(1)
	p := new(big.Int)
	for _, value := range rewardValues {
		f.Or(f, p.Lsh(p.And(f, p.Sub(p.Lsh(one, uint(value)), one)), uint(value))) // f |= (f & ((1 << v) - 1)) << v
	}
	return f.BitLen() - 1 // 答案为 f 的最高位，即 f 的二进制长度减去 1
}
