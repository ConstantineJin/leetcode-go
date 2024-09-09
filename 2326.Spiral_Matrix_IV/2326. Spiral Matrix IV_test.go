package main

import (
	"reflect"
	"testing"
)

func Test_spiralMatrix(t *testing.T) {
	type args struct {
		m    int
		n    int
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example1",
			args: args{
				m:    3,
				n:    5,
				head: &ListNode{Val: 3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8, Next: &ListNode{Val: 1, Next: &ListNode{Val: 7, Next: &ListNode{Val: 9, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 5, Next: &ListNode{Val: 0}}}}}}}}}}}}},
			},
			want: [][]int{{3, 0, 2, 6, 8}, {5, 0, -1, -1, 1}, {5, 2, 4, 9, 7}},
		},
		{
			name: "Example2",
			args: args{
				m:    1,
				n:    4,
				head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			},
			want: [][]int{{0, 1, 2, -1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralMatrix(tt.args.m, tt.args.n, tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
