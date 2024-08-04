package main

import "testing"

func Test_isSubtree(t *testing.T) {
	type args struct {
		root    *TreeNode
		subRoot *TreeNode
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
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				subRoot: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 0,
							},
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				subRoot: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
