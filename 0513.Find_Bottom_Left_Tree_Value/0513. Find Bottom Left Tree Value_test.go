package main

import "testing"

func Test_findBottomLeftValue(t *testing.T) {
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
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			wantAns: 1,
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 5,
							Left: &TreeNode{
								Val: 7,
							},
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			wantAns: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findBottomLeftValue(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("findBottomLeftValue() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
