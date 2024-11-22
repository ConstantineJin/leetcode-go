package main

import "testing"

func Test_maxEqualRowsAfterFlips(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{matrix: [][]int{{0, 1}, {1, 1}}},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{matrix: [][]int{{0, 1}, {1, 0}}},
			wantAns: 2,
		},
		{
			name:    "Example3",
			args:    args{matrix: [][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}}},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxEqualRowsAfterFlips(tt.args.matrix); gotAns != tt.wantAns {
				t.Errorf("maxEqualRowsAfterFlips() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
