package main

const mod, mx int = 1e9 + 7, 1e5 + 1

var f, g = [mx]int{1, 1, 2, 4}, [mx]int{1, 1, 2, 4}

func init() {
	for i := 4; i < mx; i++ {
		f[i] = (f[i-1] + f[i-2] + f[i-3]) % mod
		g[i] = (g[i-1] + g[i-2] + g[i-3] + g[i-4]) % mod
	}
}

func countTexts(pressedKeys string) int {
	ans, cnt := 1, 0
	for i, c := range pressedKeys {
		cnt++
		if i == len(pressedKeys)-1 || byte(c) != pressedKeys[i+1] {
			if c != '7' && c != '9' {
				ans = ans * f[cnt] % mod
			} else {
				ans = ans * g[cnt] % mod
			}
			cnt = 0
		}
	}
	return ans
}
