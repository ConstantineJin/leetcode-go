package main

import "testing"

func Test_minimumDistance(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{points: [][]int{{3, 10}, {5, 15}, {10, 2}, {4, 4}}},
			want: 12,
		},
		{
			name: "Example2",
			args: args{points: [][]int{{1, 1}, {1, 1}, {1, 1}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDistance(tt.args.points); got != tt.want {
				t.Errorf("minimumDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
