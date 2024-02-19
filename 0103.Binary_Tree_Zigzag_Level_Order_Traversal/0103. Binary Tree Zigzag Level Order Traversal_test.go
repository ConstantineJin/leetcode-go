package main

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			wantAns: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name:    "Example2",
			args:    args{root: &TreeNode{Val: 1}},
			wantAns: [][]int{{1}},
		},
		{
			name:    "Example3",
			args:    args{root: nil},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
