package main

import "testing"

func Test_sumOfPowers(t *testing.T) {
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
				nums: []int{1, 2, 3, 4},
				k:    3,
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{2, 2},
				k:    2,
			},
			want: 0,
		},
		{
			name: "Example3",
			args: args{
				nums: []int{4, 3, -1},
				k:    2,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfPowers(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("sumOfPowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
