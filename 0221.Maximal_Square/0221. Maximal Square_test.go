package main

import "testing"

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{matrix: [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{matrix: [][]byte{{'0', '1'}, {'1', '0'}}},
			wantAns: 1,
		},
		{
			name:    "Example3",
			args:    args{matrix: [][]byte{{'0'}}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximalSquare(tt.args.matrix); gotAns != tt.wantAns {
				t.Errorf("maximalSquare() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
