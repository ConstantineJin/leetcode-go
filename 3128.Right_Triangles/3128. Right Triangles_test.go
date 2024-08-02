package main

import "testing"

func Test_numberOfRightTriangles(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name:    "Example1",
			args:    args{grid: [][]int{{0, 1, 0}, {0, 1, 1}, {0, 1, 0}}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{grid: [][]int{{1, 0, 0, 0}, {0, 1, 0, 1}, {1, 0, 0, 0}}},
			wantAns: 0,
		},
		{
			name:    "Example3",
			args:    args{grid: [][]int{{1, 0, 1}, {1, 0, 0}, {1, 0, 0}}},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfRightTriangles(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("numberOfRightTriangles() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
