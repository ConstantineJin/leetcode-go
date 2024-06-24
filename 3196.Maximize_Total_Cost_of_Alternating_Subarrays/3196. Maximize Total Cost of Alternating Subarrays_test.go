package main

import "testing"

func Test_maximumTotalCost(t *testing.T) {
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
			args: args{nums: []int{1, -2, 3, 4}},
			want: 10,
		},
		{
			name: "Example2",
			args: args{nums: []int{1, -1, 1, -1}},
			want: 4,
		},
		{
			name: "Example3",
			args: args{nums: []int{0}},
			want: 0,
		},
		{
			name: "Example4",
			args: args{nums: []int{1, -1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTotalCost(tt.args.nums); got != tt.want {
				t.Errorf("maximumTotalCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
