package main

import "testing"

func Test_numberOfBoomerangs(t *testing.T) {
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
			args:    args{points: [][]int{{0, 0}, {1, 0}, {2, 0}}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{points: [][]int{{1, 1}, {2, 2}, {3, 3}}},
			wantAns: 2,
		},
		{
			name:    "Example3",
			args:    args{points: [][]int{{1, 1}}},
			wantAns: 0,
		},
		{
			name:    "Example4",
			args:    args{points: [][]int{{0, 0}, {1, 0}, {1, 1}, {0, 1}}},
			wantAns: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfBoomerangs(tt.args.points); gotAns != tt.wantAns {
				t.Errorf("numberOfBoomerangs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
