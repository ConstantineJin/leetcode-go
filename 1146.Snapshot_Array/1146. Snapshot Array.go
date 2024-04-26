package main

import "sort"

type pair struct{ snapId, val int }

type SnapshotArray struct {
	snapCnt int            // 总共进行快照的次数
	history map[int][]pair // key为原数组index, []pair为修改记录数组
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{history: make(map[int][]pair, length)}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.history[index] = append(this.history[index], pair{snapId: this.snapCnt, val: val}) // 添加修改记录
}

func (this *SnapshotArray) Snap() int {
	defer func() { this.snapCnt++ }()
	return this.snapCnt
}

func (this *SnapshotArray) Get(index int, snapId int) int {
	var h = this.history[index]
	var i = sort.Search(len(h), func(i int) bool { return h[i].snapId > snapId }) - 1 // 寻找快照编号<=snapId的最后一次修改记录等价于寻找>snapId的第一次记录-1
	if i < 0 {                                                                        // 未修改过
		return 0
	}
	return h[i].val
}
