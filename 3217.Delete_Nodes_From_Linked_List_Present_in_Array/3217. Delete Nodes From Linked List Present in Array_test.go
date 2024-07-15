package main

import (
	"reflect"
	"testing"
)

func Test_modifiedList(t *testing.T) {
	type args struct {
		nums []int
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 2, 3},
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			},
			want: &ListNode{Val: 4, Next: &ListNode{Val: 5}},
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1},
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}}}},
			},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}},
		},
		{
			name: "Example3",
			args: args{
				nums: []int{5},
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifiedList(tt.args.nums, tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifiedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
