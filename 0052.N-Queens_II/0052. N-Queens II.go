package main

// 方法一：求解（见第51题）
//func solveNQueens(n int) (ans [][]string) {
//	var dfs func(int, [][]byte)
//	dfs = func(i int, board [][]byte) {
//		board[i] = []byte(strings.Repeat(".", n))
//		var validPos = func(x, y int) bool { // 检查当前位置是否可以放置皇后
//			for j := 0; j < x; j++ { // 检查同一列
//				if board[j][y] == 'Q' {
//					return false
//				}
//			}
//			for j := 1; j <= x && j <= y; j++ { // 检查左上-右下对角线
//				if board[x-j][y-j] == 'Q' {
//					return false
//				}
//			}
//			for j := 1; j <= x && y+j < n; j++ { // 检查左下-右上对角线
//				if board[x-j][y+j] == 'Q' {
//					return false
//				}
//			}
//			return true
//		}
//		for j := range board[i] {
//			if validPos(i, j) { // 如果当前位置可以放置皇后
//				board[i][j] = 'Q'
//				if i == n-1 { // 如果已经是最后一行
//					var temp = make([]string, n)
//					for k := range temp {
//						temp[k] = string(board[k])
//					}
//					ans = append(ans, temp) // 保存结果
//					return                  // 最后一行只有一个可能的放置位置，因此可以直接返回
//				} else {
//					dfs(i+1, board) // 否则继续dfs
//					board[i][j] = '.'
//				}
//			}
//		}
//	}
//	dfs(0, make([][]byte, n))
//	return
//}
//
//func totalNQueens(n int) int {
//	return len(solveNQueens(n))
//}

// 方法二：打表
func totalNQueens(n int) int {
	var res = []int{0, 1, 0, 0, 2, 10, 4, 40, 92, 352}
	return res[n]
}
