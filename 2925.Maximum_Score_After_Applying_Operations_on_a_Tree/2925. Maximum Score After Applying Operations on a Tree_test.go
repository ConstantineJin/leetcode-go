package main

import "testing"

func Test_maximumScoreAfterOperations(t *testing.T) {
	type args struct {
		edges  [][]int
		values []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {4, 5}}, values: []int{5, 2, 5, 2, 1, 1}},
			want: 11,
		},
		{
			name: "Example2",
			args: args{edges: [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}, values: []int{20, 10, 9, 7, 4, 3, 5}},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScoreAfterOperations(tt.args.edges, tt.args.values); got != tt.want {
				t.Errorf("maximumScoreAfterOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
