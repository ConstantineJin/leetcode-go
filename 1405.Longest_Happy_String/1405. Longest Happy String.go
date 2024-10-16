package main

import "container/heap"

func longestDiverseString(a, b, c int) string {
	var ans []byte
	h := &hp{}
	if a > 0 { // 保证入堆的元素个数都大于 0，下同
		heap.Push(h, pair{'a', a})
	}
	if b > 0 {
		heap.Push(h, pair{'b', b})
	}
	if c > 0 {
		heap.Push(h, pair{'c', c})
	}
	heap.Init(h) // 堆化
	for h.Len() > 0 {
		top := heap.Pop(h).(pair)
		if len(ans) >= 2 && top.char == ans[len(ans)-1] && top.char == ans[len(ans)-2] { // 出现连续三个相同的字母（非法情况）
			if h.Len() > 0 { // 如果堆不为空
				next := heap.Pop(h).(pair) // 弹出下一个元素
				ans = append(ans, next.char)
				next.count--
				if next.count > 0 {
					heap.Push(h, next)
				}
				heap.Push(h, top)
			} else {
				break
			}
		} else { // 直接追加
			ans = append(ans, top.char)
			top.count--
			if top.count > 0 {
				heap.Push(h, top)
			}
		}
	}
	return string(ans)
}

type pair struct {
	char  byte
	count int
}

type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].count > h[j].count } // 大顶堆
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
