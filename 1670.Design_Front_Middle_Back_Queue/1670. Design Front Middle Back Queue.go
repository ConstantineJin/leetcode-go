package main

type FrontMiddleBackQueue struct {
	queue []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{queue: make([]int, 0)}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.queue = append([]int{val}, this.queue...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	var length = len(this.queue)
	this.queue = append(this.queue[:length/2], append([]int{val}, this.queue[length/2:]...)...)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.queue = append(this.queue, val)
}

func (this *FrontMiddleBackQueue) PopFront() (front int) {
	if len(this.queue) == 0 {
		return -1
	}
	front = this.queue[0]
	this.queue = this.queue[1:]
	return
}

func (this *FrontMiddleBackQueue) PopMiddle() (middle int) {
	if len(this.queue) == 0 {
		return -1
	}
	middle = this.queue[(len(this.queue)-1)/2]
	this.queue = append(this.queue[:(len(this.queue)-1)/2], this.queue[(len(this.queue)+1)/2:]...)
	return
}

func (this *FrontMiddleBackQueue) PopBack() (back int) {
	if len(this.queue) == 0 {
		return -1
	}
	back = this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]
	return
}
