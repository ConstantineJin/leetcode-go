package main

import (
	"math"
)

type MinStack struct { // 维护两个栈，stack储存每次入栈的元素，minStack储存每次入栈时栈的最小值
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, minStack: []int{math.MaxInt}}
}

func (this *MinStack) Push(val int) {
	this.stack, this.minStack = append(this.stack, val), append(this.minStack, min(val, this.GetMin())) // minStack新的栈顶是原来栈顶与此次入栈元素的较小值
}

func (this *MinStack) Pop() {
	this.stack, this.minStack = this.stack[:len(this.stack)-1], this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
