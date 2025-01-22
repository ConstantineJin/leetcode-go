package main

import "testing"

func Test_maxCoins(t *testing.T) {
	type args struct {
		piles []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{piles: []int{2, 4, 1, 2, 7, 8}},
			wantAns: 9,
		},
		{
			name:    "Example2",
			args:    args{piles: []int{2, 4, 5}},
			wantAns: 4,
		},
		{
			name:    "Example3",
			args:    args{piles: []int{9, 8, 7, 6, 5, 1, 2, 3, 4}},
			wantAns: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxCoins(tt.args.piles); gotAns != tt.wantAns {
				t.Errorf("maxCoins() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
