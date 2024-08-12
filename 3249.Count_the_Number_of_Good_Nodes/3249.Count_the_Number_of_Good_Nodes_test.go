package main

import "testing"

func Test_countGoodNodes(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{edges: [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}},
			wantAns: 7,
		},
		{
			name:    "Example2",
			args:    args{edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {0, 5}, {1, 6}, {2, 7}, {3, 8}}},
			wantAns: 6,
		},
		{
			name:    "Example3",
			args:    args{edges: [][]int{{0, 1}, {1, 2}, {1, 3}, {1, 4}, {0, 5}, {5, 6}, {6, 7}, {7, 8}, {0, 9}, {9, 10}, {9, 12}, {10, 11}}},
			wantAns: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countGoodNodes(tt.args.edges); gotAns != tt.wantAns {
				t.Errorf("countGoodNodes() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
