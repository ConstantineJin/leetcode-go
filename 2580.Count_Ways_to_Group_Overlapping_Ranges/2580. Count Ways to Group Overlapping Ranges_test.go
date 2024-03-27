package main

import "testing"

func Test_countWays(t *testing.T) {
	type args struct {
		ranges [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{ranges: [][]int{{6, 10}, {5, 15}}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{ranges: [][]int{{1, 3}, {10, 20}, {2, 5}, {4, 8}}},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countWays(tt.args.ranges); gotAns != tt.wantAns {
				t.Errorf("countWays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
