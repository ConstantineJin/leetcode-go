package main

import "testing"

func Test_sumOfMultiples(t *testing.T) {
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
			args: args{n: 7},
			want: 21,
		},
		{
			name: "Example2",
			args: args{n: 10},
			want: 40,
		},
		{
			name: "Example3",
			args: args{n: 9},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfMultiples(tt.args.n); got != tt.want {
				t.Errorf("sumOfMultiples() = %v, want %v", got, tt.want)
			}
		})
	}
}
