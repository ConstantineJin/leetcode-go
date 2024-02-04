package main

import "testing"

func Test_stoneGameVII(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{stones: []int{5, 3, 1, 4, 2}},
			want: 6,
		},
		{
			name: "Example2",
			args: args{stones: []int{7, 90, 5, 1, 100, 10, 10, 2}},
			want: 122,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameVII(tt.args.stones); got != tt.want {
				t.Errorf("stoneGameVII() = %v, want %v", got, tt.want)
			}
		})
	}
}
