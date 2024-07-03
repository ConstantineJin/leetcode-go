package main

const MX = 1001

// 方法1: 动态规划
//func minSteps(n int) int {
//	var dfs func(i, j int) (res int) // dfs(i, j) 表示当前记事本上有 i 个字符，剪贴板上有 j 个字符时，所需的最小操作数
//	dfs = func(i, j int) (res int) {
//		if i == n {
//			return 0
//		} else if i > n {
//			return MX
//		}
//		res = MX
//		if j > 0 {
//			res = min(res, dfs(i+j, j)+1) // 执行粘贴操作
//		}
//		res = min(res, dfs(i*2, i)+2) // 执行全选复制+粘贴操作
//		return
//	}
//	return dfs(1, 0) // 初始状态下记事本上只有一个字符，剪贴板为空
//}

// 方法2: 分解质因数
func minSteps(n int) (ans int) {
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			n /= i
			ans += i
		}
	}
	if n > 1 {
		ans += n
	}
	return
}
