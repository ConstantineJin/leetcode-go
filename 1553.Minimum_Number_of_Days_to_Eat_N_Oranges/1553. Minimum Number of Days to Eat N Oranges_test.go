package main

import "testing"

func Test_minDays(t *testing.T) {
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
			args: args{n: 10},
			want: 4,
		},
		{
			name: "Example2",
			args: args{n: 6},
			want: 3,
		},
		{
			name: "Example3",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "Example4",
			args: args{n: 56},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDays(tt.args.n); got != tt.want {
				t.Errorf("minDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
