package main

import "testing"

func Test_maxAncestorDiff(t *testing.T) {
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
					Val: 8,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 6,
							Left: &TreeNode{
								Val: 4,
							},
							Right: &TreeNode{
								Val: 7,
							},
						},
					},
					Right: &TreeNode{
						Val: 10,
						Right: &TreeNode{
							Val: 14,
							Left: &TreeNode{
								Val: 13,
							},
						},
					},
				},
			},
			wantAns: 7,
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 0,
							Left: &TreeNode{
								Val: 3,
							},
						},
					},
				},
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxAncestorDiff(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("maxAncestorDiff() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
