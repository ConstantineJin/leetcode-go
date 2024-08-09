package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() (_ Codec) { return }

// serialize a tree to a single string.
func (Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

// deserialize your encoded data to tree.
func (Codec) deserialize(data string) *TreeNode {
	values := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if values[0] == "null" {
			values = values[1:]
			return nil
		}
		val, _ := strconv.Atoi(values[0])
		values = values[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}
