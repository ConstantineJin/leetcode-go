package main

import "strconv"

//func maximumSwap(num int) int {
//	var arr []int
//	for x := num; x > 0; x /= 10 {
//		arr = append(arr, x%10)
//	}
//	var n = len(arr)
//	if n == 1 || slices.IsSorted(arr) {
//		return num
//	}
//	slices.Reverse(arr)
//	var sortedArr = make([]int, n)
//	copy(sortedArr, arr)
//	slices.Sort(sortedArr)
//	slices.Reverse(sortedArr)
//	var idx0, idx1, mx, ans int
//	for i, v := range arr {
//		if v != sortedArr[i] {
//			idx0 = i
//			break
//		}
//	}
//	for i := idx0; i < n; i++ {
//		if arr[i] >= mx {
//			mx, idx1 = arr[i], i
//		}
//	}
//	arr[idx0], arr[idx1] = arr[idx1], arr[idx0]
//	for i := 0; i < n; i++ {
//		ans += arr[i] * int(math.Pow10(n-i-1))
//	}
//	return ans
//}

// 一次遍历
func maximumSwap(num int) int {
	var s = strconv.Itoa(num)
	var n = len(s)
	var maxIdx = n - 1
	p, q := -1, 0
	for i := n - 2; i >= 0; i-- {
		if s[i] > s[maxIdx] { // s[i] 是目前最大数字
			maxIdx = i
		} else if s[i] < s[maxIdx] { // s[i] 右边有比它大的
			p, q = i, maxIdx // 更新 p 和 q
		}
	}
	if p == -1 { // 这意味着 s 是降序的
		return num
	}
	var t = []byte(s)
	t[p], t[q] = t[q], t[p]
	ans, _ := strconv.Atoi(string(t))
	return ans
}
