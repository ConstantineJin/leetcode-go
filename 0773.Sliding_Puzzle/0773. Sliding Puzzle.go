package main

var cache = map[string]int{"123450": 0}

func init() {
	neighbors := [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}} // 每个位置相邻的位置编号
	type pair struct {
		status      string
		index, step int // index 存储 0 所在的位置
	}
	q := []pair{{"123450", 5, 0}}
	for len(q) > 0 { // BFS 预处理所有可能产生的合法情况
		cur := q[0]
		q = q[1:]
		s := []byte(cur.status)
		x := cur.index
		for _, y := range neighbors[x] {
			s[x], s[y] = s[y], s[x]
			nxt := string(s)
			if _, ok := cache[nxt]; !ok {
				cache[nxt] = cur.step + 1
				q = append(q, pair{nxt, y, cur.step + 1})
			}
			s[x], s[y] = s[y], s[x]
		}
	}
}

func slidingPuzzle(board [][]int) int {
	s := make([]byte, 6)
	for i, row := range board {
		for j, v := range row {
			s[i*3+j] = '0' + byte(v)
		}
	}
	if step, ok := cache[string(s)]; ok {
		return step
	}
	return -1
}
