package main

type CustomStack struct {
	stack   []int
	maxSize int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{stack: make([]int, 0, maxSize), maxSize: maxSize}
}

func (s *CustomStack) Push(x int) {
	if len(s.stack) < s.maxSize {
		s.stack = append(s.stack, x)
	}
}

func (s *CustomStack) Pop() (v int) {
	n := len(s.stack)
	if n == 0 {
		return -1
	}
	v = s.stack[n-1]
	s.stack = s.stack[:n-1]
	return
}

func (s *CustomStack) Increment(k int, val int) {
	for k = min(k, len(s.stack)) - 1; k >= 0; k-- {
		s.stack[k] += val
	}
}
