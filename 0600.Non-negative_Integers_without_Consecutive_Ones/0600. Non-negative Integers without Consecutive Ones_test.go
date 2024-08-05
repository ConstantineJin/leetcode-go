package main

import "testing"

func Test_findIntegers(t *testing.T) {
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
			args: args{n: 5},
			want: 5,
		},
		{
			name: "Example2",
			args: args{n: 1},
			want: 2,
		},
		{
			name: "Example3",
			args: args{n: 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIntegers(tt.args.n); got != tt.want {
				t.Errorf("findIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
