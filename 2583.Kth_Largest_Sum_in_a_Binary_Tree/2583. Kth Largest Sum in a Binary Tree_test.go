package main

import "testing"

func Test_kthLargestLevelSum(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 4,
							},
							Right: &TreeNode{
								Val: 6,
							},
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 9,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				k: 2,
			},
			want: 13,
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
				k:    1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestLevelSum(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargestLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
