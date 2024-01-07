package main

import (
	"reflect"
	"testing"
)

func Test_insertGreatestCommonDivisors(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example1",
			args: args{head: &ListNode{Val: 18, Next: &ListNode{Val: 6, Next: &ListNode{Val: 10, Next: &ListNode{Val: 3}}}}},
			want: &ListNode{Val: 18, Next: &ListNode{Val: 6, Next: &ListNode{Val: 6, Next: &ListNode{Val: 2, Next: &ListNode{Val: 10, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}}}},
		},
		{
			name: "Example2",
			args: args{head: &ListNode{Val: 7}},
			want: &ListNode{Val: 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertGreatestCommonDivisors(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertGreatestCommonDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
