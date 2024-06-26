package main

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		cardPoints []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				cardPoints: []int{1, 2, 3, 4, 5, 6, 1},
				k:          3,
			},
			want: 12,
		},
		{
			name: "Example2",
			args: args{
				cardPoints: []int{2, 2, 2},
				k:          2,
			},
			want: 4,
		},
		{
			name: "Example3",
			args: args{
				cardPoints: []int{9, 7, 7, 9, 7, 7, 9},
				k:          7,
			},
			want: 55,
		},
		{
			name: "Example4",
			args: args{
				cardPoints: []int{1, 1000, 1},
				k:          1,
			},
			want: 1,
		},
		{
			name: "Example5",
			args: args{
				cardPoints: []int{1, 79, 80, 1, 1, 1, 200, 1},
				k:          3,
			},
			want: 202,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
