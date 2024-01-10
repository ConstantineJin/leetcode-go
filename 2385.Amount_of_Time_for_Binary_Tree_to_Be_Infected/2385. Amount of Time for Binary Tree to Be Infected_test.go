package main

import "testing"

func Test_amountOfTime(t *testing.T) {
	type args struct {
		root  *TreeNode
		start int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 9,
							},
							Right: &TreeNode{
								Val: 2,
							},
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 10,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
				start: 3,
			},
			wantAns: 4,
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
				start: 1,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := amountOfTime(tt.args.root, tt.args.start); gotAns != tt.wantAns {
				t.Errorf("amountOfTime() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
