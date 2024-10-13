package main

import "testing"

func Test_kthLargestPerfectSubtree(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 5,
							Left: &TreeNode{
								Val: 1,
							},
							Right: &TreeNode{
								Val: 8,
							},
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 5,
							Left: &TreeNode{
								Val: 6,
							},
							Right: &TreeNode{
								Val: 8,
							},
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				k: 2,
			},
			want: 3,
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
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				k: 1,
			},
			want: 7,
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
				k: 3,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestPerfectSubtree(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargestPerfectSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
