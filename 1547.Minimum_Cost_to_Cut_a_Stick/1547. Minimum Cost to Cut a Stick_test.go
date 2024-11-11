package main

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		n    int
		cuts []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				n:    7,
				cuts: []int{1, 3, 4, 5},
			},
			want: 16,
		},
		{
			name: "Example2",
			args: args{
				n:    9,
				cuts: []int{5, 6, 1, 4, 2},
			},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.n, tt.args.cuts); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
