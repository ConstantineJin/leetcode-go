package main

import (
	"reflect"
	"testing"
)

func Test_treeQueries(t *testing.T) {
	type args struct {
		root    *TreeNode
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 7,
							},
						},
					},
				},
				queries: []int{4},
			},
			want: []int{2},
		},
		{
			name: "Example2",
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
				queries: []int{3, 2, 4, 8},
			},
			want: []int{3, 2, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeQueries(tt.args.root, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
