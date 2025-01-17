package main

func doesValidArrayExist(derived []int) bool {
	var xor int
	for _, v := range derived {
		xor ^= v
	}
	return xor == 0
}
