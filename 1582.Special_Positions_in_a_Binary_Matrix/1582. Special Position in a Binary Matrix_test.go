package main

import "testing"

func Test_numSpecial(t *testing.T) {
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
			args:    args{mat: [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{mat: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numSpecial(tt.args.mat); gotAns != tt.wantAns {
				t.Errorf("numSpecial() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
