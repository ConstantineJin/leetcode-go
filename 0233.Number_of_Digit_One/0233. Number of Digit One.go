package main

func countDigitOne(n int) (ans int) {
	for mulk := 1; n >= mulk; mulk *= 10 { // k 表示从右往左第 k 位，k=0 表示个位，k=1 表示十位，以此类推。mulk 表示 10^k，即当前位的位权
		// 对于第 k 位，当前位权为 10^k(mulk)，最后 k 个数位每 10^(k+1) 个数循环一次，其中包含了 10^k 个 1，
		// 由于 n 包含了 n/(10^(k+1)) 个完整的循环，这部分共有 n/(10^(k+1))*10^k 个 1。
		// 对于不在完整循环中的 n%(10^(k+1)) 个数，其中 1 有 n%(10^(k+1))-10^k+1 个，且不会小于 0 个或大于 10^k 个
		ans += (n/(mulk*10))*mulk + min(max(n%(mulk*10)-mulk+1, 0), mulk)
	}
	return
}
