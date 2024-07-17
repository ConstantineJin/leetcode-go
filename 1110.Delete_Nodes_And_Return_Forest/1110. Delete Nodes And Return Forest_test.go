package main

import (
	"reflect"
	"testing"
)

func Test_delNodes(t *testing.T) {
	type args struct {
		root     *TreeNode
		toDelete []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []*TreeNode
	}{
		{
			name: "Example1",
			args: args{
				root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}},
				toDelete: []int{3, 5},
			},
			wantAns: []*TreeNode{{Val: 6}, {Val: 7}, {Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}}},
		},
		{
			name: "Example2",
			args: args{
				root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4}},
				toDelete: []int{3},
			},
			wantAns: []*TreeNode{{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := delNodes(tt.args.root, tt.args.toDelete); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("delNodes() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
