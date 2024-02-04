package main

import "testing"

func Test_maxSumAfterPartitioning(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				arr: []int{1, 15, 7, 9, 2, 5, 10},
				k:   3,
			},
			want: 84,
		},
		{
			name: "Example2",
			args: args{
				arr: []int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3},
				k:   4,
			},
			want: 83,
		},
		{
			name: "Example3",
			args: args{
				arr: []int{1},
				k:   1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumAfterPartitioning(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("maxSumAfterPartitioning() = %v, want %v", got, tt.want)
			}
		})
	}
}
