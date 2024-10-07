package main

import "testing"

func Test_minRefuelStops(t *testing.T) {
	type args struct {
		target    int
		startFuel int
		stations  [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				target:    1,
				startFuel: 1,
				stations:  nil,
			},
			wantAns: 0,
		},
		{
			name: "Example2",
			args: args{
				target:    100,
				startFuel: 1,
				stations:  [][]int{{10, 100}},
			},
			wantAns: -1,
		},
		{
			name: "Example3",
			args: args{
				target:    100,
				startFuel: 10,
				stations:  [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minRefuelStops(tt.args.target, tt.args.startFuel, tt.args.stations); gotAns != tt.wantAns {
				t.Errorf("minRefuelStops() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
