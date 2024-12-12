package main

import "testing"

func Test_maxSpending(t *testing.T) {
	type args struct {
		values [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{values: [][]int{{8, 5, 2}, {6, 4, 1}, {9, 7, 3}}},
			want: 285,
		},
		{
			name: "Example2",
			args: args{values: [][]int{{10, 8, 6, 4, 2}, {9, 7, 5, 3, 2}}},
			want: 386,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSpending(tt.args.values); got != tt.want {
				t.Errorf("maxSpending() = %v, want %v", got, tt.want)
			}
		})
	}
}
