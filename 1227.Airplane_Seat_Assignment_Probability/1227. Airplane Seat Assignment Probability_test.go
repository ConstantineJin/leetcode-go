package main

import "testing"

func Test_nthPersonGetsNthSeat(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Example1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "Example2",
			args: args{n: 2},
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthPersonGetsNthSeat(tt.args.n); got != tt.want {
				t.Errorf("nthPersonGetsNthSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
