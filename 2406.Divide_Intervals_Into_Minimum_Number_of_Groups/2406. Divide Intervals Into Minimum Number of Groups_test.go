package main

import "testing"

func Test_minGroups(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{intervals: [][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{intervals: [][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}}},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minGroups(tt.args.intervals); gotAns != tt.wantAns {
				t.Errorf("minGroups() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
