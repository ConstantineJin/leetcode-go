package main

import "testing"

func Test_minimumFuelCost(t *testing.T) {
	type args struct {
		roads [][]int
		seats int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "Example1",
			args: args{
				roads: [][]int{{0, 1}, {0, 2}, {0, 3}},
				seats: 5,
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				roads: [][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}},
				seats: 2,
			},
			wantAns: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumFuelCost(tt.args.roads, tt.args.seats); gotAns != tt.wantAns {
				t.Errorf("minimumFuelCost() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
