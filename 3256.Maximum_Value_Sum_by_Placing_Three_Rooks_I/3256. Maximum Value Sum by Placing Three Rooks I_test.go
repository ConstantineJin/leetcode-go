package main

import "testing"

func Test_maximumValueSum(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{board: [][]int{{-3, 1, 1, 1}, {-3, 1, -3, 1}, {-3, 2, 1, 1}}},
			want: 4,
		},
		{
			name: "Example2",
			args: args{board: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			want: 15,
		},
		{
			name: "Example3",
			args: args{board: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumValueSum(tt.args.board); got != tt.want {
				t.Errorf("maximumValueSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
