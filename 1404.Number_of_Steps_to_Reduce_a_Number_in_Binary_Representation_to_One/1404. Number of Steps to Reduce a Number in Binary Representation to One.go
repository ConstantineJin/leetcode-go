package main

func numSteps(s string) (ans int) {
	var meetOne bool // 从右向左扫描s的过程中是否遇到过'1'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			if meetOne { // 如果在右侧遇到过'1'，那么这个'0'一定会因为某一次进位变成'1'
				ans += 2 // 一次+1和一次/2
			} else { // 如果在右侧没有遇到过'1'
				ans++ // 只需要/2
			}
		} else {
			if meetOne { // 如果在右侧遇到过'1'，那么这个'1'一定会因为某一次进位变成'0'
				ans++ // 只需要/2
			} else {
				if i != 0 { // 最左侧的'1'不需要操作
					ans += 2 // 一次+1和一次/2
				}
				meetOne = true
			}
		}
	}
	return
}
