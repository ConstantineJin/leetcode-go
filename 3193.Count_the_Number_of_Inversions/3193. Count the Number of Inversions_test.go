package main

import "testing"

func Test_numberOfPermutations(t *testing.T) {
	type args struct {
		n            int
		requirements [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				n:            3,
				requirements: [][]int{{2, 2}, {0, 0}},
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				n:            3,
				requirements: [][]int{{2, 2}, {1, 1}, {0, 0}},
			},
			want: 1,
		},
		{
			name: "Example3",
			args: args{
				n:            2,
				requirements: [][]int{{0, 0}, {1, 0}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPermutations(tt.args.n, tt.args.requirements); got != tt.want {
				t.Errorf("numberOfPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
