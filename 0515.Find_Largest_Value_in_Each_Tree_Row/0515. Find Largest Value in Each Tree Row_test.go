package main

import (
	"reflect"
	"testing"
)

func Test_largestValues(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 9},
					},
				},
			},
			wantAns: []int{1, 3, 9},
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
			},
			wantAns: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := largestValues(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("largestValues() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
