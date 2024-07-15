package main

import "testing"

func Test_minimumCost(t *testing.T) {
	type args struct {
		m             int
		n             int
		horizontalCut []int
		verticalCut   []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				m:             3,
				n:             2,
				horizontalCut: []int{1, 3},
				verticalCut:   []int{5},
			},
			wantAns: 13,
		},
		{
			name: "Example2",
			args: args{
				m:             2,
				n:             2,
				horizontalCut: []int{7},
				verticalCut:   []int{4},
			},
			wantAns: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumCost(tt.args.m, tt.args.n, tt.args.horizontalCut, tt.args.verticalCut); gotAns != tt.wantAns {
				t.Errorf("minimumCost() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
