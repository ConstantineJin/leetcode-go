package main

import "strconv"

// 调用math/big包
//func addBinary(a string, b string) string {
//	// SetString方法的第一个参数为字符串，第二个参数为进制
//	ai, _ := new(big.Int).SetString(a, 2)
//	bi, _ := new(big.Int).SetString(b, 2)
//	ai.Add(ai, bi)
//	return ai.Text(2)
//}

// 直接模拟二进制相加
func addBinary(a string, b string) (sum string) {
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)
	var carry int
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		sum = strconv.Itoa(carry%2) + sum
		carry /= 2
	}
	if carry > 0 {
		sum = "1" + sum
	}
	return
}
