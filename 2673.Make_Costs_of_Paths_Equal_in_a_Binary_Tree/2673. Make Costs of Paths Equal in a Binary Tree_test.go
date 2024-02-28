package main

import "testing"

func Test_minIncrements(t *testing.T) {
	type args struct {
		n    int
		cost []int
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
				cost: []int{1, 5, 2, 2, 3, 3, 1},
			},
			want: 6,
		},
		{
			name: "Example2",
			args: args{
				n:    3,
				cost: []int{5, 3, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minIncrements(tt.args.n, tt.args.cost); got != tt.want {
				t.Errorf("minIncrements() = %v, want %v", got, tt.want)
			}
		})
	}
}
