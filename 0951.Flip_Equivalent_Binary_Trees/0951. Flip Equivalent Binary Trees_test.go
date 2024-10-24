package main

import "testing"

func Test_flipEquiv(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				root1: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{
								Val: 8,
							},
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 6,
						},
					},
				},
				root2: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 6,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
							Left: &TreeNode{
								Val: 8,
							},
							Right: &TreeNode{
								Val: 7,
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
				root1: nil,
				root2: nil,
			},
			want: true,
		},
		{
			name: "Example3",
			args: args{
				root1: nil,
				root2: &TreeNode{Val: 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipEquiv(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("flipEquiv() = %v, want %v", got, tt.want)
			}
		})
	}
}
