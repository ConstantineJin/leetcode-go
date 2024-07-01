package main

import "testing"

func Test_maximalPathQuality(t *testing.T) {
	type args struct {
		values  []int
		edges   [][]int
		maxTime int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				values:  []int{0, 32, 10, 43},
				edges:   [][]int{{0, 1, 10}, {1, 2, 15}, {0, 3, 10}},
				maxTime: 49,
			},
			wantAns: 75,
		},
		{
			name: "Example2",
			args: args{
				values:  []int{5, 10, 15, 20},
				edges:   [][]int{{0, 1, 10}, {1, 2, 10}, {0, 3, 10}},
				maxTime: 30,
			},
			wantAns: 25,
		},
		{
			name: "Example3",
			args: args{
				values:  []int{1, 2, 3, 4},
				edges:   [][]int{{0, 1, 10}, {1, 2, 11}, {2, 3, 12}, {1, 3, 13}},
				maxTime: 50,
			},
			wantAns: 7,
		},
		{
			name: "Example4",
			args: args{
				values:  []int{0, 1, 2},
				edges:   [][]int{{1, 2, 10}},
				maxTime: 10,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximalPathQuality(tt.args.values, tt.args.edges, tt.args.maxTime); gotAns != tt.wantAns {
				t.Errorf("maximalPathQuality() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
