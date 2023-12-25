package main

import "testing"

func Test_minimumCost(t *testing.T) {
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
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 6,
		},
		{
			name: "Example2",
			args: args{nums: []int{10, 12, 13, 14, 15}},
			want: 11,
		},
		{
			name: "Example3",
			args: args{nums: []int{22, 33, 22, 33, 22}},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCost(tt.args.nums); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
