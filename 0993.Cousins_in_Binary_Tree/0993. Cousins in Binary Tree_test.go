package main

import "testing"

func Test_isCousins(t *testing.T) {
	type args struct {
		root *TreeNode
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
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
					},
				},
				x: 4,
				y: 3,
			},
			want: false,
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
				x: 5,
				y: 4,
			},
			want: true,
		},
		{
			name: "Example3",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				x: 2,
				y: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCousins(tt.args.root, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}
