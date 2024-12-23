package main

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 6},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:  8,
							Left: &TreeNode{Val: 9},
						},
						Right: &TreeNode{
							Val:  5,
							Left: &TreeNode{Val: 10},
						},
					},
				},
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 6},
					},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 4},
					},
				},
			},
			wantAns: 3,
		},
		{
			name: "Example3",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 6},
					},
				},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumOperations(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("minimumOperations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
