package main

import "testing"

func Test_temperatureTrend(t *testing.T) {
	type args struct {
		temperatureA []int
		temperatureB []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				temperatureA: []int{21, 18, 18, 18, 31},
				temperatureB: []int{34, 32, 16, 16, 17},
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				temperatureA: []int{5, 10, 16, -6, 15, 11, 3},
				temperatureB: []int{16, 22, 23, 23, 25, 3, -16},
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := temperatureTrend(tt.args.temperatureA, tt.args.temperatureB); gotAns != tt.wantAns {
				t.Errorf("temperatureTrend() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
