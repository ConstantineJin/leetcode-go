package main

import "testing"

func Test_knightProbability(t *testing.T) {
	type args struct {
		n      int
		k      int
		row    int
		column int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Example1",
			args: args{
				n:      3,
				k:      2,
				row:    0,
				column: 0,
			},
			want: 0.0625,
		},
		{
			name: "Example2",
			args: args{
				n:      1,
				k:      0,
				row:    0,
				column: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knightProbability(tt.args.n, tt.args.k, tt.args.row, tt.args.column); got != tt.want {
				t.Errorf("knightProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
