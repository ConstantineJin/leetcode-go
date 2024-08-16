package main

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		arrays [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{arrays: [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{arrays: [][]int{{1}, {1}}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxDistance(tt.args.arrays); gotAns != tt.wantAns {
				t.Errorf("maxDistance() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
