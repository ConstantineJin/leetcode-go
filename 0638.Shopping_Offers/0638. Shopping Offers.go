package main

const offset = 4 // 每种物品最多购买 10 个，使用 4 位二进制数记录每种物品购买的数量

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)
	var mask int
	for i, need := range needs {
		mask |= need << (i * offset)
	}
	mem := make(map[int]int)
	var dfs func(cur int) (res int)
	dfs = func(cur int) (res int) {
		if v, ok := mem[cur]; ok {
			return v
		}
		for i := range n {
			res += price[i] * ((cur >> (i * offset)) & 0xf)
		}
		for _, s := range special {
			nxt, ok := cur, true
			for j := range n {
				if ((cur >> (j * offset)) & 0xf) < s[j] {
					ok = false
					break
				}
				nxt -= s[j] << (j * offset)
			}
			if ok {
				res = min(res, s[n]+dfs(nxt))
			}
		}
		mem[cur] = res
		return
	}
	return dfs(mask)
}
