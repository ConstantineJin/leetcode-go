package main

import "testing"

func Test_numberOfSets(t *testing.T) {
	type args struct {
		n           int
		maxDistance int
		roads       [][]int
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
				maxDistance: 5,
				roads:       [][]int{{0, 1, 20}, {0, 1, 10}, {1, 2, 2}, {0, 2, 2}},
			},
			wantAns: 7,
		},
		{
			name: "Example2",
			args: args{
				n:           1,
				maxDistance: 10,
				roads:       nil,
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfSets(tt.args.n, tt.args.maxDistance, tt.args.roads); gotAns != tt.wantAns {
				t.Errorf("numberOfSets() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
