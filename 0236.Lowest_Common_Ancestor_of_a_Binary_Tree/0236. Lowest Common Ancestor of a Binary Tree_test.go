package main

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns *TreeNode
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{
								Val: 4,
							},
						},
					},
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
				p: &TreeNode{Val: 5},
				q: &TreeNode{Val: 1},
			},
			wantAns: &TreeNode{Val: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
