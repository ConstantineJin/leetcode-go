package main

type MyCircularDeque struct {
	front, rear int
	deque       []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{deque: make([]int, k+1)}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}
	d.front = (d.front - 1 + len(d.deque)) % len(d.deque)
	d.deque[d.front] = value
	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}
	d.deque[d.rear] = value
	d.rear = (d.rear + 1) % len(d.deque)
	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}
	d.front = (d.front + 1) % len(d.deque)
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}
	d.rear = (d.rear - 1 + len(d.deque)) % len(d.deque)
	return true
}

func (d *MyCircularDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.deque[d.front]
}

func (d *MyCircularDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.deque[(d.rear-1+len(d.deque))%len(d.deque)]
}

func (d *MyCircularDeque) IsEmpty() bool { return d.front == d.rear }

func (d *MyCircularDeque) IsFull() bool { return (d.rear+1)%len(d.deque) == d.front }
