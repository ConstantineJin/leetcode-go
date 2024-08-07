package main

import "strings"

var (
	singles   = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	teens     = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens      = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	thousands = []string{"", "Thousand", "Million", "Billion"}
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	var arr []int
	for num > 0 {
		v := num % 10
		arr = append([]int{v}, arr...)
		num /= 10
	}
	n := len(arr)
	var ans []string
	switch n % 3 {
	case 1:
		ans = append(ans, singles[arr[0]])
		ans = append(ans, thousands[n/3])
	case 2:
		if arr[0] == 1 {
			ans = append(ans, teens[arr[1]])
		} else {
			ans = append(ans, tens[arr[0]])
			if arr[1] != 0 {
				ans = append(ans, singles[arr[1]])
			}
		}
		ans = append(ans, thousands[n/3])
	}
	for i := n % 3; i < n; i += 3 { // 从低位到高位每三位一起处理
		if arr[i] != 0 {
			ans = append(ans, singles[arr[i]]+" Hundred")
		}
		switch arr[i+1] {
		case 0:
			if arr[i+2] != 0 {
				ans = append(ans, singles[arr[i+2]])
			}
		case 1:
			ans = append(ans, teens[arr[i+2]])
		default:
			ans = append(ans, tens[arr[i+1]])
			if arr[i+2] != 0 {
				ans = append(ans, singles[arr[i+2]])
			}
		}
		if arr[i]+arr[i+1]+arr[i+2] > 0 {
			ans = append(ans, thousands[(n-i-2)/3])
		}
	}
	return strings.TrimSpace(strings.Join(ans, " "))
}
