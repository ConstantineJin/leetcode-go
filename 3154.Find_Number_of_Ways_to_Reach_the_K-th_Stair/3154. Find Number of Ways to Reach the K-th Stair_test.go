package main

import "testing"

func Test_waysToReachStair(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{k: 0},
			want: 2,
		},
		{
			name: "Example2",
			args: args{k: 1},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToReachStair(tt.args.k); got != tt.want {
				t.Errorf("waysToReachStair() = %v, want %v", got, tt.want)
			}
		})
	}
}
