package main

import "testing"

func Test_minCameraCover(t *testing.T) {
	type args struct {
		root *TreeNode
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
					Left: &TreeNode{
						Left:  &TreeNode{},
						Right: &TreeNode{},
					},
				},
			},
			want: 1,
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Left: &TreeNode{
						Left: &TreeNode{
							Left: &TreeNode{
								Right: &TreeNode{},
							},
						},
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCameraCover(tt.args.root); got != tt.want {
				t.Errorf("minCameraCover() = %v, want %v", got, tt.want)
			}
		})
	}
}
