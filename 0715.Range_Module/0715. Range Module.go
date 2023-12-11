package main

// RangeModule define
type RangeModule struct {
	Root *SegmentTreeNode
}

// SegmentTreeNode define
type SegmentTreeNode struct {
	Start, End  int
	Tracked     bool
	Lazy        int
	Left, Right *SegmentTreeNode
}

// Constructor define
func Constructor() RangeModule {
	return RangeModule{&SegmentTreeNode{0, 1e9, false, 0, nil, nil}}
}

// AddRange define
func (rm *RangeModule) AddRange(left int, right int) {
	update(rm.Root, left, right-1, true)
}

// QueryRange define
func (rm *RangeModule) QueryRange(left int, right int) bool {
	return query(rm.Root, left, right-1)
}

// RemoveRange define
func (rm *RangeModule) RemoveRange(left int, right int) {
	update(rm.Root, left, right-1, false)
}

func lazyUpdate(node *SegmentTreeNode) {
	if node.Lazy != 0 {
		node.Tracked = node.Lazy == 2
	}
	if node.Start != node.End {
		if node.Left == nil || node.Right == nil {
			m := node.Start + (node.End-node.Start)/2
			node.Left = &SegmentTreeNode{node.Start, m, node.Tracked, 0, nil, nil}
			node.Right = &SegmentTreeNode{m + 1, node.End, node.Tracked, 0, nil, nil}
		} else if node.Lazy != 0 {
			node.Left.Lazy = node.Lazy
			node.Right.Lazy = node.Lazy
		}
	}
	node.Lazy = 0
}

func update(node *SegmentTreeNode, start, end int, track bool) {
	lazyUpdate(node)
	if start > end || node == nil || end < node.Start || node.End < start {
		return
	}
	if start <= node.Start && node.End <= end {
		// segment completely covered by the update range
		node.Tracked = track
		if node.Start != node.End {
			if track {
				node.Left.Lazy = 2
				node.Right.Lazy = 2
			} else {
				node.Left.Lazy = 1
				node.Right.Lazy = 1
			}
		}
		return
	}
	update(node.Left, start, end, track)
	update(node.Right, start, end, track)
	node.Tracked = node.Left.Tracked && node.Right.Tracked
}

func query(node *SegmentTreeNode, start, end int) bool {
	lazyUpdate(node)
	if start > end || node == nil || end < node.Start || node.End < start {
		return true
	}
	if start <= node.Start && node.End <= end {
		// segment completely covered by the update range
		return node.Tracked
	}
	return query(node.Left, start, end) && query(node.Right, start, end)
}
