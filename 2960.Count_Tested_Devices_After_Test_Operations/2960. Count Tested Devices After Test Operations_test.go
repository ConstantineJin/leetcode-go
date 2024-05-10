package main

import "testing"

func Test_countTestedDevices(t *testing.T) {
	type args struct {
		batteryPercentages []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{batteryPercentages: []int{1, 1, 2, 1, 3}},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{batteryPercentages: []int{0, 1, 2}},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countTestedDevices(tt.args.batteryPercentages); gotAns != tt.wantAns {
				t.Errorf("countTestedDevices() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
