package main

import "testing"

func Test_lastRemaining(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{n: 9},
			want: 6,
		},
		{
			name: "Example2",
			args: args{n: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastRemaining(tt.args.n); got != tt.want {
				t.Errorf("lastRemaining() = %v, want %v", got, tt.want)
			}
		})
	}
}
