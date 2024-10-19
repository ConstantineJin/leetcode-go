package main

func findKthBit(n, k int) byte {
	if k == 1 {
		return '0'
	}
	mid := 1 << (n - 1)
	if k == mid {
		return '1'
	} else if k < mid {
		return findKthBit(n-1, k)
	} else {
		k = mid*2 - k
		if findKthBit(n-1, k) == '1' {
			return '0'
		} else {
			return '1'
		}
	}
}
