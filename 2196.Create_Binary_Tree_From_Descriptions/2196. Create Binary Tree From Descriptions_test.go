package main

import (
	"reflect"
	"testing"
)

func Test_createBinaryTree(t *testing.T) {
	type args struct {
		descriptions [][]int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example1",
			args: args{descriptions: [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}},
			want: &TreeNode{Val: 50, Left: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 17}}, Right: &TreeNode{Val: 80, Left: &TreeNode{Val: 19}}},
		},
		{
			name: "Example2",
			args: args{descriptions: [][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}}},
			want: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createBinaryTree(tt.args.descriptions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
