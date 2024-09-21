package main

import "testing"

func Test_countSpecialNumbers(t *testing.T) {
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
			args: args{n: 20},
			want: 19,
		},
		{
			name: "Example2",
			args: args{n: 5},
			want: 5,
		},
		{
			name: "Example3",
			args: args{n: 135},
			want: 110,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSpecialNumbers(tt.args.n); got != tt.want {
				t.Errorf("countSpecialNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
