package main

import "testing"

func Test_maxScoreSightseeingPair(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{values: []int{8, 1, 5, 2, 6}},
			wantAns: 11,
		},
		{
			name:    "Example2",
			args:    args{values: []int{1, 2}},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxScoreSightseeingPair(tt.args.values); gotAns != tt.wantAns {
				t.Errorf("maxScoreSightseeingPair() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
