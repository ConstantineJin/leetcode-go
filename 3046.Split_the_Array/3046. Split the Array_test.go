package main

import "testing"

func Test_isPossibleToSplit(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{nums: []int{1, 1, 2, 2, 3, 4}},
			want: true,
		},
		{
			name: "Example2",
			args: args{nums: []int{1, 1, 1, 1}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossibleToSplit(tt.args.nums); got != tt.want {
				t.Errorf("isPossibleToSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
