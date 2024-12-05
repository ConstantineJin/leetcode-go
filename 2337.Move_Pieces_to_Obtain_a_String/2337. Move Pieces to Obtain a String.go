package main

import "strings"

func canChange(start, target string) bool {
	if strings.ReplaceAll(start, "_", "") != strings.ReplaceAll(target, "_", "") { // 'L' 和 'R' 都无法“穿透”对方，如果去掉所有 '_' 后的字符串不相同直接返回 false
		return false
	}
	var j int // 双指针
	for i, c := range start {
		if c != '_' {
			for target[j] == '_' { // 通过所有空格
				j++
			}
			for i != j && c == 'L' == (i < j) { // 如果当前字符为 'L' 且 i<j，由于 'L' 无法向右移动，返回 false。反之亦然
				return false
			}
			j++
		}
	}
	return true
}
