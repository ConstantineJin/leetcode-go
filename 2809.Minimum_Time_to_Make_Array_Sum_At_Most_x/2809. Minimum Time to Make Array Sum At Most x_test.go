package main

import "testing"

func Test_minimumTime(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		x     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{1, 2, 3},
				x:     4,
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{3, 3, 3},
				x:     4,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTime(tt.args.nums1, tt.args.nums2, tt.args.x); got != tt.want {
				t.Errorf("minimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
