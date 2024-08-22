package main

import "testing"

func Test_findMaximumNumber(t *testing.T) {
	type args struct {
		k int64
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
				k: 9,
				x: 1,
			},
			want: 6,
		},
		{
			name: "Example2",
			args: args{
				k: 7,
				x: 2,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumNumber(tt.args.k, tt.args.x); got != tt.want {
				t.Errorf("findMaximumNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
