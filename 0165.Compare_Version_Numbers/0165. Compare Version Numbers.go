package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	ver1, ver2, i := strings.Split(version1, "."), strings.Split(version2, "."), 0 // ver1和ver2分别记录version1和version2按照"."分割所得的[]string，i记录当前遍历到的下标
	len1, len2 := len(ver1), len(ver2)
	for ; i < min(len1, len2); i++ { // 保证此时两个版本号都有对应段
		v1, _ := strconv.Atoi(ver1[i])
		v2, _ := strconv.Atoi(ver2[i])
		if v1 > v2 { // 如果某个版本号在该段比另一个大，则直接返回
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	if len1 > len2 {
		for ; i < len1; i++ {
			if v1, _ := strconv.Atoi(ver1[i]); v1 != 0 { // 只要多出来的版本号有一个不为0，则其版本更高
				return 1
			}
		}
		return 0 // 多出来的版本号全是0，两者版本相等
	} else if len1 < len2 {
		for ; i < len2; i++ {
			if v2, _ := strconv.Atoi(ver2[i]); v2 != 0 {
				return -1
			}
		}
		return 0
	} else { // 长度相同，则版本号相同
		return 0
	}
}
