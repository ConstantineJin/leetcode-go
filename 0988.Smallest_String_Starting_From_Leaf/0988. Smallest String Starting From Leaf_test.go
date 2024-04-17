package main

import "testing"

func Test_smallestFromLeaf(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
			},
			want: "dba",
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val: 25,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
			},
			want: "adz",
		},
		{
			name: "Example3",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 1,
							Left: &TreeNode{
								Val: 1,
							},
						},
					},
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 0,
						},
					},
				},
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestFromLeaf(tt.args.root); got != tt.want {
				t.Errorf("smallestFromLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
