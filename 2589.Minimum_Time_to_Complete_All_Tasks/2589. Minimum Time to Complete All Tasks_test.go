package main

import "testing"

func Test_findMinimumTime(t *testing.T) {
	type args struct {
		tasks [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{tasks: [][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{tasks: [][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}}},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findMinimumTime(tt.args.tasks); gotAns != tt.wantAns {
				t.Errorf("findMinimumTime() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
