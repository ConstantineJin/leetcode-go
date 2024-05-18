package main

import "testing"

func Test_distributeCoins(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 0}}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{root: &TreeNode{Val: 0, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 0}}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := distributeCoins(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("distributeCoins() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
