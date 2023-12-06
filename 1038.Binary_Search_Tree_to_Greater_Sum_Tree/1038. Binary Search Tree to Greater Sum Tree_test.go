package main

import (
	"reflect"
	"testing"
)

func Test_bstToGst(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 2,
							Right: &TreeNode{
								Val: 3,
							},
						},
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 7,
							Right: &TreeNode{
								Val: 8,
							},
						},
					},
				},
			},
			want: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val: 36,
					Left: &TreeNode{
						Val: 36,
					},
					Right: &TreeNode{
						Val: 35,
						Right: &TreeNode{
							Val: 33,
						},
					},
				},
				Right: &TreeNode{
					Val: 21,
					Left: &TreeNode{
						Val: 26,
					},
					Right: &TreeNode{
						Val: 15,
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val: 0,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstToGst(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstToGst() = %v, want %v", got, tt.want)
			}
		})
	}
}
