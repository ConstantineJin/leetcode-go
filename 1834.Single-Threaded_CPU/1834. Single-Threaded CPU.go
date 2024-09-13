package main

import (
	"container/heap"
	"sort"
)

type Task struct{ id, enqueueTime, processingTime int }

func getOrder(_tasks [][]int) []int {
	n := len(_tasks)
	tasks := make([]Task, 0, n)
	for i, task := range _tasks {
		tasks = append(tasks, Task{i, task[0], task[1]})
	}
	sort.Slice(tasks, func(i, j int) bool { return tasks[i].enqueueTime < tasks[j].enqueueTime }) // 按照任务入队时间排序
	t := tasks[0].enqueueTime                                                                     // 系统时间
	h := &hp{}                                                                                    // 使用堆维护正在等待运行的任务
	var i int                                                                                     // 标记哪些任务已经入堆
	ans := make([]int, 0, n)
	for i < n || h.Len() > 0 { // 还有任务没有入堆或者在堆中
		for i < n && tasks[i].enqueueTime <= t { // 将当前任务执行期间入队的任务入堆
			heap.Push(h, tasks[i])
			i++
		}
		if h.Len() == 0 { // 如果堆为空，直接跳转到下一个任务入队的时间
			t = tasks[i].enqueueTime
			continue
		}
		task := heap.Pop(h).(Task) // 出堆
		ans = append(ans, task.id)
		t += task.processingTime
	}
	return ans
}

type hp []Task

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	return h[i].processingTime < h[j].processingTime || h[i].processingTime == h[j].processingTime && h[i].id < h[j].id
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(Task)) }
func (h *hp) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
