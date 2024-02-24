package main

import (
	"reflect"
	"testing"
)

func Test_closestNodes(t *testing.T) {
	type args struct {
		root    *TreeNode
		queries []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 13,
						Left: &TreeNode{
							Val: 9,
						},
						Right: &TreeNode{
							Val: 15,
							Left: &TreeNode{
								Val: 14,
							},
						},
					},
				},
				queries: []int{2, 5, 16},
			},
			want: [][]int{{2, 2}, {4, 6}, {15, -1}},
		},
		{
			name: "Example2",
			args: args{
				root:    &TreeNode{Val: 4, Right: &TreeNode{Val: 9}},
				queries: []int{3},
			},
			want: [][]int{{-1, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestNodes(tt.args.root, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
