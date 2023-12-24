package main

import "testing"

func Test_minStoneSum(t *testing.T) {
	type args struct {
		piles []int
		k     int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				piles: []int{5, 4, 9},
				k:     2,
			},
			wantAns: 12,
		},
		{
			name: "Example2",
			args: args{
				piles: []int{4, 3, 6, 7},
				k:     3,
			},
			wantAns: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minStoneSum(tt.args.piles, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("minStoneSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
