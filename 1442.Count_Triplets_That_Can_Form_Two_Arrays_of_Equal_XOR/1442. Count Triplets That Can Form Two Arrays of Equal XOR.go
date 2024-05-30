package main

// 引理1:两个连续相邻的子数组arr[i:j]和arr[j:k+1]的异或和相等，则两个子数组合并的子数组arr[i:k+1]的异或和必为0，且与j无关
// 引理2:若子数组arr[i:k+1]的异或和为0，则i和k+1的前缀异或和相等
// 方法1: 双重循环，时间复杂度O(n^2)
//func countTriplets(arr []int) (ans int) {
//	n := len(arr)
//	prefixXORs := make([]int, n+1) // 前缀异或和
//	for i, v := range arr {
//		prefixXORs[i+1] = prefixXORs[i] ^ v
//	}
//	// 通过双重循环遍历子数组的异或和，如果子数组arr[i:k+1]的异或和为0，那么j可以选择(i,k]区间内的任意数，共有k-i个可能性
//	for i := 0; i < n; i++ {
//		for k := i + 1; k < n; k++ {
//			if prefixXORs[i] == prefixXORs[k+1] {
//				ans += k - i
//			}
//		}
//	}
//	return
//}

// 方法2: 哈希表，时间复杂度O(n)
func countTriplets(arr []int) (ans int) {
	cnt, total := make(map[int]int), make(map[int]int) // cnt用于存储前缀异或和的出现次数，total用于存储前缀异或和的下标之和
	var prefixXOR int
	for k, v := range arr {
		if m, ok := cnt[prefixXOR^v]; ok {
			ans += m*k - total[prefixXOR^v]
		}
		cnt[prefixXOR]++
		total[prefixXOR] += k
		prefixXOR ^= v
	}
	return
}
