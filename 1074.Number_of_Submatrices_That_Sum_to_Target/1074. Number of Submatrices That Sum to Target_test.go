package main

import "testing"

func Test_numSubmatrixSumTarget(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				matrix: [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}},
				target: 0,
			},
			wantAns: 4,
		},
		{
			name: "Example2",
			args: args{
				matrix: [][]int{{1, -1}, {-1, 1}},
				target: 0,
			},
			wantAns: 5,
		},
		{
			name: "Example3",
			args: args{
				matrix: [][]int{{904}},
				target: 0,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numSubmatrixSumTarget(tt.args.matrix, tt.args.target); gotAns != tt.wantAns {
				t.Errorf("numSubmatrixSumTarget() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
