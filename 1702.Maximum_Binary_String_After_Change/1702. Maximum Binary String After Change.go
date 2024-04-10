package main

import "strings"

// 贪心策略：结果字符串中最多只含一个0，否则可通过操作2将后续的0不断左移至第一个0处再使用操作1进行替换
func maximumBinaryString(binary string) string {
	var i = strings.Index(binary, "0")
	if i < 0 { // 原字符串全为1，已是最大
		return binary
	}
	var cnt1 = strings.Count(binary[i:], "1")
	return strings.Repeat("1", len(binary)-cnt1-1) + "0" + strings.Repeat("1", cnt1) // 结果字符串中唯一的0的下标是n-cnt1-1
}
