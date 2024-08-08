package main

import "testing"

func Test_minimumAddedInteger(t *testing.T) {
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
				nums1: []int{4, 20, 16, 12, 8},
				nums2: []int{14, 18, 10},
			},
			want: -2,
		},
		{
			name: "Example2",
			args: args{
				nums1: []int{3, 5, 5, 3},
				nums2: []int{7, 7},
			},
			want: 2,
		},
		{
			name: "Example3",
			args: args{
				nums1: []int{10, 2, 8, 7, 5, 6, 7, 10},
				nums2: []int{5, 8, 5, 3, 8, 4},
			},
			want: -2,
		},
		{
			name: "Example4",
			args: args{
				nums1: []int{4, 6, 3, 1, 4, 2, 10, 9, 5},
				nums2: []int{5, 10, 3, 2, 6, 1, 9},
			},
			want: 0,
		},
		{
			name: "Example5",
			args: args{
				nums1: []int{7, 2, 6, 8, 7},
				nums2: []int{7, 6, 5},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAddedInteger(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minimumAddedInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
