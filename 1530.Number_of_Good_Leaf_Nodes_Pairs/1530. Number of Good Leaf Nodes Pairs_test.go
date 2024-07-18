package main

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		root     *TreeNode
		distance int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
				distance: 3,
			},
			wantAns: 1,
		},
		{
			name: "Example2",
			args: args{
				root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}},
				distance: 3,
			},
			wantAns: 2,
		},
		{
			name: "Example3",
			args: args{
				root:     &TreeNode{Val: 7, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}}},
				distance: 3,
			},
			wantAns: 1,
		},
		{
			name: "Example4",
			args: args{
				root:     &TreeNode{Val: 100},
				distance: 1,
			},
			wantAns: 0,
		},
		{
			name: "Example5",
			args: args{
				root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}},
				distance: 2,
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countPairs(tt.args.root, tt.args.distance); gotAns != tt.wantAns {
				t.Errorf("countPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
