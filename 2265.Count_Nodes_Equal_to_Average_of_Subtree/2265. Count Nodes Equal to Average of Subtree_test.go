package main

import "testing"

func Test_averageOfSubtree(t *testing.T) {
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
					Val: 4,
					Left: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			wantAns: 5,
		},
		{
			name:    "Example2",
			args:    args{root: &TreeNode{Val: 1}},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := averageOfSubtree(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("averageOfSubtree() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
