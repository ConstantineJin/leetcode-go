package main

import "testing"

func Test_countCompleteDayPairs(t *testing.T) {
	type args struct {
		hours []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{hours: []int{12, 12, 30, 24, 24}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{hours: []int{72, 48, 24, 3}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countCompleteDayPairs(tt.args.hours); gotAns != tt.wantAns {
				t.Errorf("countCompleteDayPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
