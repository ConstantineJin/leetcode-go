package main

import "testing"

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{grid: [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{grid: [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numIslands(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("numIslands() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
