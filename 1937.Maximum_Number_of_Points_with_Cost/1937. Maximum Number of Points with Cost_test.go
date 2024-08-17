package main

import "testing"

func Test_maxPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{points: [][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}},
			want: 9,
		},
		{
			name: "Example2",
			args: args{points: [][]int{{1, 5}, {2, 3}, {4, 2}}},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
