package main

import (
	path2 "path"
)

// 不使用任何库函数
//func simplifyPath(path string) (ans string) {
//	var st []string
//	var i, j int
//	var n = len(path)
//	for ; i < n; i = j + 1 {
//		for ; i < n && path[i] == '/'; i++ {
//		}
//		for j = i; j < n && path[j] != '/'; j++ {
//		}
//		switch path[i:j] {
//		case "", ".":
//		case "..":
//			if len(st) > 0 {
//				st = st[:len(st)-1]
//			}
//		default:
//			st = append(st, path[i:j])
//		}
//	}
//	if len(st) == 0 {
//		return "/"
//	}
//	for _, s := range st {
//		ans += "/" + s
//	}
//	return
//}

// 标准库
func simplifyPath(path string) string {
	return path2.Clean(path)
}
