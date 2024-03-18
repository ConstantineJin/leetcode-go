package main

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name:      "Example1",
			args:      args{points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}},
			wantCount: 2,
		},
		{
			name:      "Example2",
			args:      args{points: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}},
			wantCount: 4,
		},
		{
			name:      "Example3",
			args:      args{points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}},
			wantCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := findMinArrowShots(tt.args.points); gotCount != tt.wantCount {
				t.Errorf("findMinArrowShots() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
