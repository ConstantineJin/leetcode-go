package main

import "testing"

func Test_maxPossibleScore(t *testing.T) {
	type args struct {
		start []int
		d     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				start: []int{6, 0, 3},
				d:     2,
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				start: []int{2, 6, 13, 13},
				d:     5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPossibleScore(tt.args.start, tt.args.d); got != tt.want {
				t.Errorf("maxPossibleScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
