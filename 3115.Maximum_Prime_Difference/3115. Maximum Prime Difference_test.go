package main

import "testing"

func Test_maximumPrimeDifference(t *testing.T) {
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
			args: args{nums: []int{4, 2, 9, 5, 3}},
			want: 3,
		},
		{
			name: "Example2",
			args: args{nums: []int{4, 8, 2, 8}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumPrimeDifference(tt.args.nums); got != tt.want {
				t.Errorf("maximumPrimeDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
