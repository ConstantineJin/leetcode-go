package main

import "sort"

const mx = 1000

var primes = []int{0}

func init() {
	np := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func primeSubOperation(nums []int) bool {
	var pre int // 前一个数减完质数后的结果
	for _, num := range nums {
		if num <= pre { // 无法满足单调性
			return false
		}
		pre = num - primes[sort.SearchInts(primes, num-pre)-1] // 若 p 是满足 num-p > pre 的最大质数，则 p 是满足 p < num-pre 的最大质数
	}
	return true
}
