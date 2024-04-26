package main

import "testing"

func Test_minFallingPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			want: 13,
		},
		{
			name: "Example2",
			args: args{grid: [][]int{{7}}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minFallingPathSum(tt.args.grid); gotAns != tt.want {
				t.Errorf("minFallingPathSum() = %v, want %v", gotAns, tt.want)
			}
		})
	}
}
