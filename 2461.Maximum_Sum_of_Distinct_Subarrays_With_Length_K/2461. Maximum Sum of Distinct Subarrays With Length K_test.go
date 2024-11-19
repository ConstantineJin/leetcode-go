package main

import "testing"

func Test_maximumSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 5, 4, 2, 9, 9, 9},
				k:    3,
			},
			want: 15,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{4, 4, 4},
				k:    3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
