package main

import "testing"

func Test_canCross(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{stones: []int{0, 1, 3, 5, 6, 8, 12, 17}},
			want: true,
		},
		{
			name: "Example2",
			args: args{stones: []int{0, 1, 2, 3, 4, 8, 9, 11}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCross(tt.args.stones); got != tt.want {
				t.Errorf("canCross() = %v, want %v", got, tt.want)
			}
		})
	}
}
