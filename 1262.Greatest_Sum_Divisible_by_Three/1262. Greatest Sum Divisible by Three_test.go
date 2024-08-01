package main

import "testing"

func Test_maxSumDivThree(t *testing.T) {
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
			args: args{nums: []int{3, 6, 5, 1, 8}},
			want: 18,
		},
		{
			name: "Example2",
			args: args{nums: []int{4}},
			want: 0,
		},
		{
			name: "Example3",
			args: args{nums: []int{1, 2, 3, 4, 4}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumDivThree(tt.args.nums); got != tt.want {
				t.Errorf("maxSumDivThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
