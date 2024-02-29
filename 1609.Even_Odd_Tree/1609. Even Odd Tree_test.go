package main

import "testing"

func Test_isEvenOddTree(t *testing.T) {
	type args struct {
		root *TreeNode
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
						Val: 10,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 12,
							},
							Right: &TreeNode{
								Val: 8,
							},
						},
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 7,
							Left: &TreeNode{
								Val: 6,
							},
						},
						Right: &TreeNode{
							Val: 9,
							Right: &TreeNode{
								Val: 2,
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEvenOddTree(tt.args.root); got != tt.want {
				t.Errorf("isEvenOddTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
