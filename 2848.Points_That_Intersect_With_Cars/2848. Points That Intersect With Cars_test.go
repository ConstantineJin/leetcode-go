package main

import "testing"

func Test_numberOfPoints(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{nums: [][]int{{3, 6}, {1, 5}, {4, 7}}},
			wantAns: 7,
		},
		{
			name:    "Example2",
			args:    args{nums: [][]int{{1, 3}, {5, 8}}},
			wantAns: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfPoints(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("numberOfPoints() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
