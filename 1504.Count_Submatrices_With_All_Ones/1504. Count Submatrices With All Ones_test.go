package main

import "testing"

func Test_numSubmat(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{mat: [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}},
			wantAns: 13,
		},
		{
			name:    "Example2",
			args:    args{mat: [][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}}},
			wantAns: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numSubmat(tt.args.mat); gotAns != tt.wantAns {
				t.Errorf("numSubmat() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
