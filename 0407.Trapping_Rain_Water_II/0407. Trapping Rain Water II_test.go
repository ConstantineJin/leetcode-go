package main

import "testing"

func Test_trapRainWater(t *testing.T) {
	type args struct {
		heightMap [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{heightMap: [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{heightMap: [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}}},
			wantAns: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := trapRainWater(tt.args.heightMap); gotAns != tt.wantAns {
				t.Errorf("trapRainWater() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
