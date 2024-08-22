package main

import "testing"

func Test_minEnd(t *testing.T) {
	type args struct {
		n int
		x int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{
				n: 3,
				x: 4,
			},
			want: 6,
		},
		{
			name: "Example2",
			args: args{
				n: 2,
				x: 7,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEnd(tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("minEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
