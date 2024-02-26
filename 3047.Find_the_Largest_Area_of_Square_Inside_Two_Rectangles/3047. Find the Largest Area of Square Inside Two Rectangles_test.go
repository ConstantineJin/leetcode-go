package main

import "testing"

func Test_largestSquareArea(t *testing.T) {
	type args struct {
		bottomLeft [][]int
		topRight   [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example0",
			args: args{
				bottomLeft: [][]int{{2, 3}, {1, 4}},
				topRight:   [][]int{{3, 5}, {5, 5}},
			},
			want: 1,
		},
		{
			name: "Example1",
			args: args{
				bottomLeft: [][]int{{3, 2}, {1, 1}},
				topRight:   [][]int{{4, 5}, {5, 4}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSquareArea(tt.args.bottomLeft, tt.args.topRight); got != tt.want {
				t.Errorf("largestSquareArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
