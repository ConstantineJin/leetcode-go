package main

import "testing"

func Test_maxResult(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, -1, -2, 4, -7, 3},
				k:    2,
			},
			want: 7,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{10, -5, -2, 4, 0, 3},
				k:    3,
			},
			want: 17,
		},
		{
			name: "Example3",
			args: args{
				nums: []int{1, -5, -20, 4, -1, 3, -6, -3},
				k:    2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxResult(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
