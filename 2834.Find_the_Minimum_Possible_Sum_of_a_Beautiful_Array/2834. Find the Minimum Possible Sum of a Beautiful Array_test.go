package main

import "testing"

func Test_minimumPossibleSum(t *testing.T) {
	type args struct {
		n      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{n: 2, target: 3},
			want: 4,
		},
		{
			name: "Example2",
			args: args{n: 3, target: 3},
			want: 8,
		},
		{
			name: "Example3",
			args: args{n: 1, target: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPossibleSum(tt.args.n, tt.args.target); got != tt.want {
				t.Errorf("minimumPossibleSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
