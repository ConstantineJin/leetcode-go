package main

import "testing"

func Test_removeStones(t *testing.T) {
	type args struct {
		stones [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{stones: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}},
			want: 5,
		},
		{
			name: "Example2",
			args: args{stones: [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}},
			want: 3,
		},
		{
			name: "Example3",
			args: args{stones: [][]int{{0, 0}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeStones(tt.args.stones); got != tt.want {
				t.Errorf("removeStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
