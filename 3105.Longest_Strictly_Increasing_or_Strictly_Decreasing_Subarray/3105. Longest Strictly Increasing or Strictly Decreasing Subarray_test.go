package main

import "testing"

func Test_longestMonotonicSubarray(t *testing.T) {
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
			args: args{nums: []int{1, 4, 3, 3, 2}},
			want: 2,
		},
		{
			name: "Example2",
			args: args{nums: []int{3, 3, 3, 3}},
			want: 1,
		},
		{
			name: "Example3",
			args: args{nums: []int{3, 2, 1}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMonotonicSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestMonotonicSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
