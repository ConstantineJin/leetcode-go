package main

import "testing"

func Test_canReachCorner(t *testing.T) {
	type args struct {
		X       int
		Y       int
		circles [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				X:       3,
				Y:       4,
				circles: [][]int{{2, 1, 1}},
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				X:       3,
				Y:       3,
				circles: [][]int{{1, 1, 2}},
			},
			want: false,
		},
		{
			name: "Example3",
			args: args{
				X:       3,
				Y:       3,
				circles: [][]int{{2, 1, 1}, {1, 2, 1}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReachCorner(tt.args.X, tt.args.Y, tt.args.circles); got != tt.want {
				t.Errorf("canReachCorner() = %v, want %v", got, tt.want)
			}
		})
	}
}
