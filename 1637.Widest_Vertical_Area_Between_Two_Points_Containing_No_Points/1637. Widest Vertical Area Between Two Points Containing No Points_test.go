package main

import "testing"

func Test_maxWidthOfVerticalArea(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{points: [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{points: [][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxWidthOfVerticalArea(tt.args.points); gotAns != tt.wantAns {
				t.Errorf("maxWidthOfVerticalArea() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
