package main

import "testing"

func Test_nonSpecialCount(t *testing.T) {
	type args struct {
		l int
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				l: 5,
				r: 7,
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				l: 4,
				r: 16,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nonSpecialCount(tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("nonSpecialCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
