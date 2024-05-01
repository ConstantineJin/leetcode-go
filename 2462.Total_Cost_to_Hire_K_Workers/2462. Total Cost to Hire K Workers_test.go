package main

import "testing"

func Test_totalCost(t *testing.T) {
	type args struct {
		costs      []int
		k          int
		candidates int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "Example1",
			args: args{
				costs:      []int{17, 12, 10, 2, 7, 2, 11, 20, 8},
				k:          3,
				candidates: 4,
			},
			wantAns: 11,
		},
		{
			name: "Example2",
			args: args{
				costs:      []int{1, 2, 4, 1},
				k:          3,
				candidates: 3,
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := totalCost(tt.args.costs, tt.args.k, tt.args.candidates); gotAns != tt.wantAns {
				t.Errorf("totalCost() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
