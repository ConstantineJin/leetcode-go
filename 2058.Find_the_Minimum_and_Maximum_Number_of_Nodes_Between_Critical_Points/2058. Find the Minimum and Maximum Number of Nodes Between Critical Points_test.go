package main

import (
	"reflect"
	"testing"
)

func Test_nodesBetweenCriticalPoints(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{head: &ListNode{Val: 3, Next: &ListNode{Val: 1}}},
			want: []int{-1, -1},
		},
		{
			name: "Example2",
			args: args{head: &ListNode{Val: 5, Next: &ListNode{Val: 3, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}}}}}},
			want: []int{1, 3},
		},
		{
			name: "Example3",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 7}}}}}}}}}},
			want: []int{3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nodesBetweenCriticalPoints(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nodesBetweenCriticalPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
