package main

import "testing"

func Test_paintWalls(t *testing.T) {
	type args struct {
		cost []int
		time []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				cost: []int{1, 2, 3, 2},
				time: []int{1, 2, 3, 2},
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				cost: []int{2, 3, 4, 2},
				time: []int{1, 1, 1, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paintWalls(tt.args.cost, tt.args.time); got != tt.want {
				t.Errorf("paintWalls() = %v, want %v", got, tt.want)
			}
		})
	}
}
