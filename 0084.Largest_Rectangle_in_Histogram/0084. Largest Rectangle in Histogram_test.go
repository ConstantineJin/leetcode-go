package main

import "testing"

func Test_largestRectangleArea(t *testing.T) {
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
			args:    args{heights: []int{2, 1, 5, 6, 2, 3}},
			wantAns: 10,
		},
		{
			name:    "Example2",
			args:    args{heights: []int{2, 4}},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := largestRectangleArea(tt.args.heights); gotAns != tt.wantAns {
				t.Errorf("largestRectangleArea() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
