package main

import (
	"math/big"
	"slices"
	"sort"
)

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	rewardValues = slices.Compact(rewardValues)
	one := big.NewInt(1)
	f := big.NewInt(1)
	p := new(big.Int)
	for _, value := range rewardValues {
		mask := p.Sub(p.Lsh(one, uint(value)), one)
		f.Or(f, p.Lsh(p.And(f, mask), uint(value)))
	}
	return f.BitLen() - 1
}
