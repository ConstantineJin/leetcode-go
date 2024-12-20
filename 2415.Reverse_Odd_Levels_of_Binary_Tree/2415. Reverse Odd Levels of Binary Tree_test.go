package main

import (
	"reflect"
	"testing"
)

func Test_reverseOddLevels(t *testing.T) {
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
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 8,
						},
						Right: &TreeNode{
							Val: 13,
						},
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 21,
						},
						Right: &TreeNode{
							Val: 34,
						},
					},
				},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 8,
					},
					Right: &TreeNode{
						Val: 13,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 21,
					},
					Right: &TreeNode{
						Val: 34,
					},
				},
			},
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 13},
					Right: &TreeNode{Val: 11},
				},
			},
			want: &TreeNode{
				Val:   7,
				Left:  &TreeNode{Val: 11},
				Right: &TreeNode{Val: 13},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOddLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseOddLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
