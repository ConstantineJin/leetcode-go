package main

import "testing"

func Test_calPoints(t *testing.T) {
	type args struct {
		operations []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{operations: []string{"5", "2", "C", "D", "+"}},
			wantAns: 30,
		},
		{
			name:    "Example2",
			args:    args{operations: []string{"5", "-2", "4", "C", "D", "9", "+", "+"}},
			wantAns: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := calPoints(tt.args.operations); gotAns != tt.wantAns {
				t.Errorf("calPoints() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
