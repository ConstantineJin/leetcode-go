package main

import (
	"reflect"
	"testing"
)

func Test_allPossibleFBT(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []*TreeNode
	}{
		{
			name:    "Example2",
			args:    args{n: 3},
			wantAns: []*TreeNode{{Val: 0, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 0}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPossibleFBT(tt.args.n); !reflect.DeepEqual(got, tt.wantAns) {
				t.Errorf("allPossibleFBT() = %v, want %v", got, tt.wantAns)
			}
		})
	}
}
