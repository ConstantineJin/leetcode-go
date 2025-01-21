package main

import "testing"

func Test_maxValueOfCoins(t *testing.T) {
	type args struct {
		piles [][]int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				piles: [][]int{{1, 100, 3}, {7, 8, 9}},
				k:     2,
			},
			want: 101,
		},
		{
			name: "Example2",
			args: args{
				piles: [][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}},
				k:     7,
			},
			want: 706,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValueOfCoins(tt.args.piles, tt.args.k); got != tt.want {
				t.Errorf("maxValueOfCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
