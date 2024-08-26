package main

import "testing"

func Test_maxSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{nums: []int{51, 71, 17, 24, 42}},
			want: 88,
		},
		{
			name: "Example2",
			args: args{nums: []int{1, 2, 3, 4}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.args.nums); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
