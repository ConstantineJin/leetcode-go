package main

// 方法1：深度优先搜索
//type jugs struct{ jug1, jug2 int }
//
//func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
//	var vis = make(map[jugs]bool)
//	var dfs func(i, j int) bool
//	dfs = func(i, j int) bool {
//		if vis[jugs{i, j}] { // 避免无限递归，如果检查过当前状态就返回false
//			return false
//		}
//		vis[jugs{i, j}] = true
//		if i == targetCapacity || j == targetCapacity || i+j == targetCapacity { // 其中一壶或两壶之和为target，返回true
//			return true
//		}
//		if dfs(jug1Capacity, j) || dfs(i, jug2Capacity) || dfs(0, j) || dfs(i, 0) { // dfs其中一壶装满或倒空，如果可以得到target就返回true
//			return true
//		}
//		var a, b = min(i, jug2Capacity-j), min(j, jug1Capacity-i) // 将一壶的水倒入另一壶中，dfs可以得到target就返回true
//		return dfs(i-a, j+a) || dfs(i+b, j-b)
//	}
//	return dfs(0, 0)
//}

// 方法2：裴蜀定理
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	return jug1Capacity+jug2Capacity >= targetCapacity && targetCapacity%gcd(jug1Capacity, jug2Capacity) == 0
}
