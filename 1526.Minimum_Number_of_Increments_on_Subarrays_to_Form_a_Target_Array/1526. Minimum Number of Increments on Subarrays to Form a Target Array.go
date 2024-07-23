package main

// 差分数组
func minNumberOperations(target []int) (ans int) {
	ans = target[0] // 对于 i=0 需要 target[0] 次操作
	for i := 1; i < len(target); i++ {
		ans += max(target[i]-target[i-1], 0) // 如果 target[i-1] < target[i]，那么对 target[i-1] 进行 +1 操作时，可以免费对 target[i] 也进行 +1 操作，这样还需额外 target[i] - target[i-1] 次操作。否则不需要额外操作次数
	}
	return
}
