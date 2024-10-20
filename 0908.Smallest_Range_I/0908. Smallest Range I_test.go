package main

import "testing"

func Test_smallestRangeI(t *testing.T) {
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
				nums: []int{1},
				k:    0,
			},
			want: 0,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{0, 10},
				k:    2,
			},
			want: 6,
		},
		{
			name: "Example3",
			args: args{
				nums: []int{1, 3, 6},
				k:    3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRangeI(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("smallestRangeI() = %v, want %v", got, tt.want)
			}
		})
	}
}
