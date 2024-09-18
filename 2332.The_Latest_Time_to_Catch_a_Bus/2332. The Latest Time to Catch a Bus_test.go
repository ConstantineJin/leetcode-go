package main

import "testing"

func Test_latestTimeCatchTheBus(t *testing.T) {
	type args struct {
		buses      []int
		passengers []int
		capacity   int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				buses:      []int{10, 20},
				passengers: []int{2, 17, 18, 19},
				capacity:   2,
			},
			wantAns: 16,
		},
		{
			name: "Example2",
			args: args{
				buses:      []int{20, 30, 10},
				passengers: []int{19, 13, 26, 4, 25, 11, 21},
				capacity:   2,
			},
			wantAns: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := latestTimeCatchTheBus(tt.args.buses, tt.args.passengers, tt.args.capacity); gotAns != tt.wantAns {
				t.Errorf("latestTimeCatchTheBus() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
