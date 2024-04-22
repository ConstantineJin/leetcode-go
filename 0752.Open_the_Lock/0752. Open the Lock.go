package main

func openLock(deadends []string, target string) (ans int) {
	if target == "0000" {
		return
	}
	var dead, visited = make(map[string]bool), make(map[string]bool)
	for _, deadend := range deadends {
		dead[deadend] = true
	}
	if dead["0000"] {
		return -1
	}
	var cur = []string{"0000"}
	for len(cur) > 0 {
		var nxt []string
		for _, s := range cur {
			if s == target {
				return ans
			}
			if dead[s] || visited[s] {
				continue
			}
			visited[s] = true
			for i := 0; i < 4; i++ {
				var temp = []byte(s)
				if temp[i] > '0' {
					temp[i]--
				} else {
					temp[i] = '9'
				}
				nxt = append(nxt, string(temp))
				temp = []byte(s)
				if temp[i] < '9' {
					temp[i]++
				} else {
					temp[i] = '0'
				}
				nxt = append(nxt, string(temp))
			}
		}
		cur = nxt
		ans++
	}
	return -1
}
