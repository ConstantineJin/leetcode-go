package main

import "testing"

func Test_maxTwoEvents(t *testing.T) {
	type args struct {
		events [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{events: [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{events: [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}},
			wantAns: 5,
		},
		{
			name:    "Example3",
			args:    args{events: [][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}},
			wantAns: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxTwoEvents(tt.args.events); gotAns != tt.wantAns {
				t.Errorf("maxTwoEvents() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
