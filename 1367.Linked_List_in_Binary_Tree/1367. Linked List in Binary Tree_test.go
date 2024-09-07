package main

import "testing"

func Test_isSubPath(t *testing.T) {
	type args struct {
		head *ListNode
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
				head: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 8}}},
				root: &TreeNode{Val: 1, Left: &TreeNode{Val: 4, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}}}},
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6}}}},
				root: &TreeNode{Val: 1, Left: &TreeNode{Val: 4, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}}}},
			},
			want: true,
		},
		{
			name: "Example3",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8}}}}},
				root: &TreeNode{Val: 1, Left: &TreeNode{Val: 4, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}}}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubPath(tt.args.head, tt.args.root); got != tt.want {
				t.Errorf("isSubPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
