package main

import "testing"

func Test_countHousePlacements(t *testing.T) {
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
			args: args{n: 1},
			want: 4,
		},
		{
			name: "Example2",
			args: args{n: 2},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHousePlacements(tt.args.n); got != tt.want {
				t.Errorf("countHousePlacements() = %v, want %v", got, tt.want)
			}
		})
	}
}
