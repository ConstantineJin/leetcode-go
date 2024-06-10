package main

import "testing"

func Test_heightChecker(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{heights: []int{1, 1, 4, 2, 1, 3}},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{heights: []int{5, 1, 2, 3, 4}},
			wantAns: 5,
		},
		{
			name:    "Example3",
			args:    args{heights: []int{1, 2, 3, 4, 5}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := heightChecker(tt.args.heights); gotAns != tt.wantAns {
				t.Errorf("heightChecker() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
