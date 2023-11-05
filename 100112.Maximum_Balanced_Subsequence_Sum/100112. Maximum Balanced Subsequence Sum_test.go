package main

import "testing"

func Test_maxBalancedSubsequenceSum(t *testing.T) {
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
			args: args{nums: []int{3, 3, 5, 6}},
			want: 14,
		},
		{
			name: "Example2",
			args: args{nums: []int{5, -1, -3, 8}},
			want: 13,
		},
		{
			name: "Example3",
			args: args{nums: []int{-2, -1}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxBalancedSubsequenceSum(tt.args.nums); got != tt.want {
				t.Errorf("maxBalancedSubsequenceSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
