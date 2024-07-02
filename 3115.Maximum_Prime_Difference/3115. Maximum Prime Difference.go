package main

const mx = 101

var notPrime = [mx]bool{true, true}

func init() { // 埃氏筛
	for i := 2; i*i < mx; i++ {
		if !notPrime[i] {
			for j := i * i; j < mx; j += i {
				notPrime[j] = true
			}
		}
	}
}

func maximumPrimeDifference(nums []int) int {
	i, j := 0, len(nums)-1
	for notPrime[nums[i]] {
		i++
	}
	for notPrime[nums[j]] {
		j--
	}
	return j - i
}
