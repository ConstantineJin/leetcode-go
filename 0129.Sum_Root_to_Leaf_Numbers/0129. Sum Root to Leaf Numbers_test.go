package main

import "testing"

func Test_sumNumbers(t *testing.T) {
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
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			wantAns: 25,
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 9,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 0,
					},
				},
			},
			wantAns: 1026,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := sumNumbers(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("sumNumbers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
