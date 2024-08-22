package main

import "math/bits"

func findComplement(num int) int {
	return num ^ (1<<bits.Len(uint(num)) - 1) // 生成一个与 num 长度相同的全 1 二进制数，然后与 num 做异或
}
