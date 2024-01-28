package main

import "testing"

func Test_maxNumberOfAlloys(t *testing.T) {
	type args struct {
		n           int
		k           int
		budget      int
		composition [][]int
		stock       []int
		cost        []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				n:           3,
				k:           2,
				budget:      15,
				composition: [][]int{{1, 1, 1}, {1, 1, 10}},
				stock:       []int{0, 0, 0},
				cost:        []int{1, 2, 3},
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				n:           3,
				k:           2,
				budget:      15,
				composition: [][]int{{1, 1, 1}, {1, 1, 10}},
				stock:       []int{0, 0, 100},
				cost:        []int{1, 2, 3},
			},
			wantAns: 5,
		},
		{
			name: "Example3",
			args: args{
				n:           2,
				k:           3,
				budget:      10,
				composition: [][]int{{2, 1}, {1, 2}, {1, 1}},
				stock:       []int{1, 1},
				cost:        []int{5, 5},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxNumberOfAlloys(tt.args.n, tt.args.k, tt.args.budget, tt.args.composition, tt.args.stock, tt.args.cost); gotAns != tt.wantAns {
				t.Errorf("maxNumberOfAlloys() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
