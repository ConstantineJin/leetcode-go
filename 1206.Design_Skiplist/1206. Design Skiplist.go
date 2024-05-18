package main

import "math/rand/v2"

const (
	maxLevel = 32
	pFactor  = 0.25
)

type node struct {
	val     int
	forward []*node
}

type Skiplist struct {
	head  *node
	level int
}

func Constructor() Skiplist {
	return Skiplist{
		head: &node{
			val:     -1,
			forward: make([]*node, maxLevel),
		},
		level: 0,
	}
}

func (s *Skiplist) randomLevel() int {
	level := 1
	for level < maxLevel && rand.Float64() < pFactor {
		level++
	}
	return level
}

// Search 返回 target 是否存在于跳表中
func (s *Skiplist) Search(target int) bool {
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层最小且最接近 target 的元素
		for cur.forward[i] != nil && cur.forward[i].val < target {
			cur = cur.forward[i]
		}
	}
	cur = cur.forward[0]
	// 检查当前元素是否就是 target
	return cur != nil && cur.val == target
}

// Add 向跳表中插入元素 num
func (s *Skiplist) Add(num int) {
	update := make([]*node, maxLevel) // update 存储每一层需要更新 forward 指针的节点
	for i := range update {
		update[i] = s.head
	}
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层最小且最接近 num 的元素
		for cur.forward[i] != nil && cur.forward[i].val < num {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	level := s.randomLevel()
	s.level = max(s.level, level)
	newNode := &node{val: num, forward: make([]*node, level)}
	// 逐层更新，将当前元素的 forward 指向插入的新节点
	for i, nd := range update[:level] {
		newNode.forward[i] = nd.forward[i]
		nd.forward[i] = newNode
	}
}

// Erase 删除跳表中的给定值 num ，如果不存在 num 返回false；如果存在多个删除任意一个
func (s *Skiplist) Erase(num int) bool {
	update := make([]*node, maxLevel)
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].val < num {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	cur = cur.forward[0]
	if cur == nil || cur.val != num { // num 不存在跳表中
		return false
	}
	for i := 0; i < s.level && update[i].forward[i] == cur; i++ {
		// 对第 i 层进行更新，将 forward 指向被删除节点的下一跳
		update[i].forward[i] = cur.forward[i]
	}
	// 更新当前 level
	for s.level > 1 && s.head.forward[s.level-1] == nil { // 最上层没有任何节点
		s.level--
	}
	return true
}
