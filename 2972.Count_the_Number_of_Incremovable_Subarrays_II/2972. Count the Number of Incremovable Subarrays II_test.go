package main

import "testing"

func Test_incremovableSubarrayCount(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 10,
		},
		{
			name: "Example2",
			args: args{nums: []int{6, 5, 7, 8}},
			want: 7,
		},
		{
			name: "Example3",
			args: args{nums: []int{8, 7, 6, 6}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incremovableSubarrayCount(tt.args.nums); got != tt.want {
				t.Errorf("incremovableSubarrayCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
