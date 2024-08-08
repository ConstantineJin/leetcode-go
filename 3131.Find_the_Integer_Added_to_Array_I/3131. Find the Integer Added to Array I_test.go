package main

import "testing"

func Test_addedInteger(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				nums1: []int{2, 6, 4},
				nums2: []int{9, 7, 5},
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				nums1: []int{10},
				nums2: []int{5},
			},
			want: -5,
		},
		{
			name: "Example3",
			args: args{
				nums1: []int{1, 1, 1, 1},
				nums2: []int{1, 1, 1, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addedInteger(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("addedInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
