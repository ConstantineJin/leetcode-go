package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{prices: []int{7, 1, 5, 3, 6, 4}},
			wantAns: 7,
		},
		{
			name:    "Example2",
			args:    args{prices: []int{1, 2, 3, 4, 5}},
			wantAns: 4,
		},
		{
			name:    "Example3",
			args:    args{prices: []int{7, 6, 4, 3, 1}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxProfit(tt.args.prices); gotAns != tt.wantAns {
				t.Errorf("maxProfit() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
