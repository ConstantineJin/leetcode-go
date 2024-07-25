package main

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{num: "2245047"},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{num: "2908305"},
			wantAns: 3,
		},
		{
			name:    "Example3",
			args:    args{num: "10"},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumOperations(tt.args.num); gotAns != tt.wantAns {
				t.Errorf("minimumOperations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
