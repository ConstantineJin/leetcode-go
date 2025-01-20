package main

import "testing"

func Test_findClosestNumber(t *testing.T) {
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
			args: args{nums: []int{-4, -2, 1, 4, 8}},
			want: 1,
		},
		{
			name: "Example2",
			args: args{nums: []int{2, -1, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestNumber(tt.args.nums); got != tt.want {
				t.Errorf("findClosestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
